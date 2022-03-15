package db

import "time"

type TWorkflow struct {
	WfId      int8      `json:"wf_id" gorm:"autoIncrement; primaryKey"`
	WftId     int8      `json:"wft_id"`
	WfName    string    `json:"wf_name"`
	WfDesc    string    `json:"wf_desc"`
	WfCount   int8      `json:"wf_count"`
	WfStatus  int8      `json:"wf_status"` //预留字段
	CreatedAt time.Time `json:"created_at"`
}

const (
	// //WfStatusDefault int8 = iota

	WfStatusDenied   int8 = -1
	WfStatusApproved int8 = 1
)

func InsertWorkflow(row *TWorkflow) error {
	return db.Model(&TWorkflow{}).Create(row).Error
}

// GetLastWorkflows 获取最新的证书申请流程信息
func GetLastWorkflows() (rows []*TWorkflow, err error) {
	err = db.Model(&TWorkflow{}).Group("wft_id").Having("max(wf_id)").Find(&rows).Error
	return
}

func GetLastWorkflow(t int8) (row TWorkflow, err error) {
	err = db.Model(&TWorkflow{}).Where("wft_id = ?", t).Last(&row).Error
	return
}

func GetWorkflowById(id int8) (row TWorkflow, err error) {
	err = db.Model(&TWorkflow{}).First(&row, id).Error
	return
}

func UpdateWfStatus(id, status int8) error {
	return db.Model(&TWorkflow{WftId: id}).Update("status", status).Error
}
