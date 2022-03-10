package db

import (
	"fmt"
	"time"
)

// TApply 申请信息表
type TApply struct {
	ApplyId   int       `json:"apply_id" gorm:"column:apply_id;primaryKey;autoIncrement"`
	WfId      int8      `json:"wf_id"`      //流程id
	WfNodeId  int       `json:"wf_node_id"` //当前待办流程节点id
	PInfoId   int       `json:"p_info_id"`  //个人证书申请信息
	EntOrgId  int       `json:"ent_org_id"`
	EntAntId  int       `json:"ent_ant_id"`
	UserId    string    `json:"user_id"`
	DeviceId  string    `json:"device_id"`
	Status    int8      `json:"status"` //申请状态
	CreatedAt time.Time `json:"created_at"`
}

func GetApply(id string) (res []*TApply, err error) {

	roleSql := fmt.Sprintf("select t_user_role.id from t_user_role where t_user_role.user_id = \"%s\"", id)
	nodeSql := fmt.Sprintf("select t_wf_node_role.wf_node_id from t_wf_node_role where t_wf_node_role.role_id in (%s)", roleSql)

	sql := fmt.Sprintf("select t_apply.* from t_apply where t_apply.wf_node_id in (%s)", nodeSql)

	err = db.Raw(sql).Scan(&res).Error
	return
}
