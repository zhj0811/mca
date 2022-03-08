package db

import "gorm.io/gorm/clause"

// TWfType workflow type table
type TWfType struct {
	WftId     int8   `gorm:"column:wft_id;primaryKey;autoIncrement"` //类型id
	WftName   string `json:"wft_name"`                               //类型名称
	WftDesc   string `json:"wft_desc"`                               //类型描述
	WftStatus int8   `json:"wft_status"`                             //类型状态
}

const (
	WfTypePersonalCertVerify int8 = iota + 1
	WfTypeEnterpriseCertVerify
)

func initWft() error {
	wfts := []TWfType{
		{WftId: WfTypePersonalCertVerify, WftName: "个人证书申请流程", WftDesc: "个人证书申请流程"},
		{WftId: WfTypeEnterpriseCertVerify, WftName: "企业证书申请流程", WftDesc: "企业证书申请流程"},
	}
	for _, row := range wfts {
		err := db.Model(&TWfType{}).Clauses(clause.OnConflict{
			DoNothing: true,
		}).Create(&row).Error
		if err != nil {
			return err
		}
	}
	return nil
}
