package db

type TWorkflow struct {
	WfId     int8   `json:"wf_id" gorm:"autoIncrement; primaryKey"`
	WfName   string `json:"wf_name"`
	WfDesc   string `json:"wf_desc"`
	WfCount  int8   `json:"wf_count"`
	WftId    int8   `json:"wft_id"`
	WfStatus int8   `json:"wf_status"` //预留字段
}

func InsertWorkflow(row *TWorkflow) error {
	return db.Model(&TWorkflow{}).Create(row).Error
}
