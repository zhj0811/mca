package db

import "fmt"

type TWfNodeRole struct {
	Id       int  `json:"id" gorm:"autoIncrement; primaryKey"` //	自动编号	主键
	WfNodeId int  `json:"wf_node_id"`
	RoleId   int8 `json:"role_id"`
	Status   int8 `json:"status"` //预留字段
}

func InsertWfNodeRole(row *TWfNodeRole) error {
	return db.Model(&TWfNodeRole{}).Create(row).Error
}

func InsertWfNodeRoles(row []TWfNodeRole) error {
	return db.Model(&TWfNodeRole{}).Create(&row).Error
}

func DeleteWfNodeRole(id int) error {
	return db.Delete(&TWfNodeRole{}, id).Error
}

type SpecWorkflow struct {
	TWfNode
	WfNodeRoleId int  `json:"wf_node_role_id"`
	RoleId       int8 `json:"role_id"`
}

func GetSpecWorkflow(id string) (res []SpecWorkflow, err error) {
	sql := fmt.Sprintf("select node.*, role.id as wf_node_role_id, role.role_id "+
		"from t_wf_node as node left join t_wf_node_role as role on node.wf_node_id = role.wf_node_id where node.wf_id = %s", id)
	err = db.Raw(sql).Scan(&res).Error
	return
}
