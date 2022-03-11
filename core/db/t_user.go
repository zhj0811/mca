package db

import (
	"fmt"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"jzsg.com/mca/core/server/config"
	"jzsg.com/mca/core/utils"
)

type TUser struct {
	UserId   string `json:"user_id" gorm:"column:user_id;primaryKey;not null"`
	Name     string `json:"name" gorm:"uniqueIndex;not null"`
	Nickname string `json:"nickname"`
	Phone    string `json:"phone"`
	Passwd   string `json:"passwd"`
	Email    string `json:"email"`
	Status   int8   `json:"status"`
}

const (
	UserStatusDisable int8 = -1
)

func InsertUser(user *TUser) error {
	if user.UserId == "" {
		user.UserId = primitive.NewObjectID().Hex()
	}
	return db.Model(&TUser{}).Create(user).Error
}

func PutUser(row *TUser) error {
	return db.Model(row).Omit("passwd").Save(row).Error //忽略密码
	//return db.Save(row).Error
}

var adminId string

func initAdmin() error {
	admin := config.GetAdminConf()
	var err error
	adminId, err = GetAdminId()
	if err != nil {
		return err
	}
	if adminId != "" {
		return nil
	}

	user := &TUser{Name: admin.User, Passwd: utils.MD5Bytes([]byte(utils.MD5Bytes([]byte(admin.Passwd))))}
	err = InsertUser(user)
	adminId = user.UserId
	err = db.Model(&TUserRole{}).Create(&TUserRole{
		RoleId: RoleAdmin,
		UserId: user.UserId,
	}).Error

	return err
}

func GetUserByName(name string) (res TUser, err error) {
	err = db.Model(&TUser{}).Where("name = ?", name).First(&res).Error
	return
}

func GetUserById(id string) (res TUser, err error) {
	err = db.Model(&TUser{}).Where("user_id = ?", id).First(&res).Error
	return
}

func GetAllOpr(name, status string, limit, offset int) (res []*TUser, err error) {
	var nameSql, statusSql string
	if name != "" {
		nameSql = fmt.Sprintf(" AND (name LIKE \"%%%s%%\" OR nickname LIKE \"%%%s%%\")", name, name)
	}

	if status != "" {
		statusSql = fmt.Sprintf(" AND status = %s", status)
	}

	sql := fmt.Sprintf("select t_user.* from t_user where user_id != \"%s\"%s%s limit %d, %d", adminId, nameSql, statusSql, offset, limit)
	err = db.Raw(sql).Scan(&res).Error
	return
}

func IsAdminRole(id string) int8 {
	if id == adminId {
		return RoleAdmin
	}
	return RoleOperator
}

func UpdatePasswd(id, passwd string) error {
	err := db.Model(&TUser{Name: id}).Update("passwd", passwd).Error
	return err
}

func IsValidName(name string) (valid bool, err error) {
	var count int64
	err = db.Model(&TUser{}).Where("name = ?", name).Count(&count).Error
	if err != nil {
		return
	}
	valid = count == 0
	return
}
