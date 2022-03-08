package db

// TApply 申请信息表
type TApply struct {
	ApplyId  int    `json:"apply_id" gorm:"column:apply_id;primaryKey;autoIncrement"`
	WfId     int8   `json:"wf_id"`      //流程id
	WfNodeId int    `json:"wf_node_id"` //当前待办流程节点id
	PInfoId  int    `json:"p_info_id"`  //个人证书申请信息
	EntOrgId int    `json:"ent_org_id"`
	EntAntId int    `json:"ent_ant_id"`
	UserId   string `json:"user_id"`
	DeviceId string `json:"device_id"`
	Status   int8   `json:"status"` //申请状态
}
