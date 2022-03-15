package db

type TWfOperation struct {
	Id       int    `json:"id" gorm:"autoIncrement; primaryKey"` //	自动编号	主键
	ApplyId  int    `json:"apply_id"`                            //申请表id
	WfNodeId int    `json:"wf_node_id"`                          //流程节点id
	Remark   string `json:"remark"`                              //备忘录
	UserId   string `json:"user_id"`                             //审核员id
	OprDesc  string `json:"opr_desc"`                            //html 前端备忘 	操作描述 同意 拒绝
}

func InsertWfOpr(row *TWfOperation) error {
	return db.Create(row).Error
}
