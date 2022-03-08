package db

import (
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

func InsertUser(user *TUser) error {
	if user.UserId == "" {
		user.UserId = primitive.NewObjectID().Hex()
	}
	return db.Model(&TUser{}).Create(user).Error
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
