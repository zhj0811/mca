package db

type TEnterApplyOrgInfo struct {
	EntOrgId      int    `json:"ent_org_id" gorm:"autoIncrement; primaryKey"`
	OrgName       string //机构姓名
	OrgIdeType    bool   //机构证件类型	0：营业执照 1 其他
	OrgIdeCard    string //机构证件号码
	OrgIdePic     string //机构证件图片链接
	LegalName     string //法人姓名
	LegalIdeType  int    //法人证件类型	0：身份证 1 其他
	LegalIdeCard  string //法人证件号码
	LegalElecSeal string //法人印章图片链接
	HashType      int    //hash算法类型
}
