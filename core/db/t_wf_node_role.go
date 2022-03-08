package db

type TWfNodeRole struct {
	Id       int  `json:"id" gorm:"autoIncrement; primaryKey"` //	自动编号	主键
	WfNodeId int  `json:"wf_node_id"`
	RoleId   int8 `json:"role_id"`
	Status   int8 `json:"status"` //预留字段
}

func InsertWfNodeRole(row *TWfNodeRole) error {
	return db.Model(&TWfNodeRole{}).Create(row).Error
}

func InsertWfNodeRoles(row *[]TWfNodeRole) error {
	return db.Model(&TWfNodeRole{}).Create(row).Error
}
