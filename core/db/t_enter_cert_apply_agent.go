package db

type TEnterApplyAgentInfo struct {
	AgntId        int    `json:"agnt_id" gorm:"column:agnt_id;primaryKey;autoIncrement"` //自动编号	主键
	AgntName      string `json:"agnt_name"`                                              //经办人真实姓名
	AgntIdeType   bool   `json:"agnt_ide_type"`                                          //经办人证件类型	0：身份证 1 其他
	AgntIdeCard   string `json:"agnt_ide_card"`                                          //经办人证件号码
	AgntIdeFront  string `json:"agnt_ide_front"`                                         //证件照正面链接
	AgntIdeBack   string `json:"agnt_ide_back"`                                          //证件照反面链接
	AgntEmail     string `json:"agnt_email"`                                             //经办人邮箱
	AgntOrg       string `json:"agnt_org"`                                               //所属机构
	AgntDuty      string `json:"agnt_duty"`                                              //职务
	Phone         string `json:"phone"`                                                  //电话
	AuzPic        string `json:"auz_pic"`                                                //授权书图片链接地址
	AgntIdeStatus int8   `json:"agnt_ide_status"`                                        //经办人证件校验状态	0 default  1成功 -1 失败
}
