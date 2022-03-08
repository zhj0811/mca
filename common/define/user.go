package define

type User struct {
	User   string `json:"user" yaml:"user"`     //用户名
	Passwd string `json:"passwd" yaml:"passwd"` //密码
}
