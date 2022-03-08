package db

// TWfNode workflow node table
type TWfNode struct {
	WfNodeId    int    `json:"wf_node_id" gorm:"autoIncrement; primaryKey"`
	WfId        int8   `json:"wf_id"`
	WfNodeName  string `json:"wf_node_name"`
	WfNodeDesc  string `json:"wf_node_desc"`
	WfNodeIndex int8   `json:"wf_node_index"`
	WfNodeType  int8   // 节点类型 0: 开始 1：结束 2：操作
}

func InsertWfNode(row *TWfNode) error {
	return db.Model(&TWfNode{}).Create(row).Error
}
