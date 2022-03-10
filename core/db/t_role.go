package db

import (
	"fmt"
	"time"

	"gorm.io/gorm/clause"
)

// TRole 用户角色表
type TRole struct {
	RoleId     int8      `json:"role_id" gorm:"column:role_id;primaryKey;autoIncrement"`
	RoleName   string    `json:"role_name" gorm:"unique; not null"` //角色名称
	RoleDesc   string    `json:"role_desc"`                         //角色描述
	RoleStatus int8      `json:"status"`                            //预留字段
	CreatedAt  time.Time `json:"created_at"`
}

const (
	RoleAdmin int8 = iota + 1
	RoleOperator
)

func InitRole() error {
	roles := []TRole{
		{RoleId: RoleAdmin, RoleName: "超管"},
		{RoleId: RoleOperator, RoleName: "操作员"},
	}
	for _, row := range roles {
		err := db.Model(&TRole{}).Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&row).Error
		if err != nil {
			return err
		}
	}
	return nil
}

func InsertRole(row *TRole) error {
	return db.Model(&TRole{}).Create(row).Error
}

func GetRoles() (rows *[]TRole, err error) {
	err = db.Where("role_id > ?", RoleOperator).Find(rows).Error
	return
}

func GetSpecUserRole(id string) (rows []*TRole, err error) {
	roleSql := fmt.Sprintf("select t_user_role.role_id from t_user_role where t_user_role.user_id = \"%s\"", id)
	sql := fmt.Sprintf("select t_role.* from t_role where t_role.role_id in (%s)", roleSql)
	err = db.Raw(sql).Scan(&rows).Error
	return
}
