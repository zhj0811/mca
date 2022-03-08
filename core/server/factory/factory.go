package factory

import "jzsg.com/mca/common/zlogging"

var logger = zlogging.MustGetLogger("handler")

type UsersForRoleReq struct {
	Role int8     `json:"role"`
	User []string `json:"user"`
}

type RolesForUserReq struct {
	User string `json:"user"`
	Role []int8 `json:"role"`
}
