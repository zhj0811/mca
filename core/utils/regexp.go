package utils

import "regexp"

const (
	RegexpTypeName int8 = iota + 1
	RegexpPhone
	RegexpEmail
)

func RegexpStr(t int8, str string) bool {
	var expr string
	switch t {
	case RegexpTypeName:
		//reg, _ = regexp.Compile("[A-Za-z0-9_-]{5,20}")
		expr = "^[A-Za-z0-9_-]{5,20}$"
	case RegexpPhone:
		expr = `1[3|5|7|8|][\d]{9}`
	case RegexpEmail:
		expr = `\w+([-+.]\w+)*@\w+([-.]\w+)*\.\w+([-.]\w+)*`
	default:
		return false
	}
	reg, _ := regexp.Compile(expr)
	return reg.MatchString(str)
}
