package db

// TPersonCertApplyInfo 个人证书申请信息表
type TPersonCertApplyInfo struct {
	PInfoId   int    `json:"p_info_id" gorm:"column:p_info_id;primaryKey;autoIncrement"`
	PName     string `json:"p_name"`
	IdeType   bool   //个人证件类型	0：身份证 1 其他
	IdeCard   string //个人证件号码
	IdeFront  string //证件照正面链接
	IdeBack   string //证件照反面链接
	Email     string //申请人邮箱
	Org       string //所属机构
	Duty      string //职务
	BankId    string //银行卡号
	ElecSeal  string //电子印章链接
	HashType  int    //hash算法类型
	Phone     string //电话
	IdeStatus int    //证件校验状态	0审核中 1成功 -1 失败
}
