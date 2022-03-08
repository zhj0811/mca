package db

type TFile struct {
	Id       int    `json:"id" gorm:"autoIncrement; primaryKey"` //自动编号
	UserId   string `json:"user_id"`                             //上传用户id
	Path     string `json:"path"`                                //存放路径
	Status   int8   `json:"status"`                              //0 正常 1删除
	FileType int8   `yaml:"file_type"`                           //文件类型	身份证 营业执照 头像
	Hash     string `yaml:"hash"`                                //文件Hash
}
