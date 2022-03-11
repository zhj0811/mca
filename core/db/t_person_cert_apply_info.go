package db

// TPersonCertApplyInfo 个人证书申请信息表
type TPersonCertApplyInfo struct {
	PInfoId   int    `json:"p_info_id" gorm:"column:p_info_id;primaryKey;autoIncrement"`
	PName     string `json:"p_name"`     //申请人姓名
	IdeType   bool   `json:"ide_type"`   //个人证件类型	0：身份证 1 其他
	IdeCard   string `json:"ide_card"`   //个人证件号码
	IdeFront  string `json:"ide_front"`  //证件照正面链接
	IdeBack   string `json:"ide_back"`   //证件照反面链接
	Email     string `json:"email"`      //申请人邮箱
	Org       string `json:"org"`        //所属机构
	Duty      string `json:"duty"`       //职务
	BankId    string `json:"bank_id"`    //银行卡号
	ElecSeal  string `json:"elec_seal"`  //电子印章链接
	HashType  int    `json:"hash_type"`  //hash算法类型
	Phone     string `json:"phone"`      //电话
	IdeStatus int    `json:"ide_status"` //证件校验状态	0审核中 1成功 -1 失败
}

func InsertIndCertApplyInfo(row *TPersonCertApplyInfo) error {
	return db.Model(&TPersonCertApplyInfo{}).Create(row).Error
}
