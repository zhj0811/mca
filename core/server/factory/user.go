package factory

import (
	"github.com/pkg/errors"
	"jzsg.com/mca/common/define"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
	"jzsg.com/mca/core/utils"
)

type loginRes struct {
	Token string `json:"token" yaml:"token"`
	Role  int8   `json:"role"`
}

func Login(req define.User) common.ResponseInfo {
	res := common.ResponseInfo{
		Code: common.Success,
	}
	//var res common.ResponseInfo
	user, err := db.GetUserByName(req.User)
	if err != nil {
		res.Code = common.UserNameOrPasswordErr
		res.Msg = err.Error()
		return res
	}

	if user.Passwd != utils.MD5Bytes([]byte(req.Passwd)) {
		res.Code = common.UserNameOrPasswordErr
		res.Msg = "password error"
		return res
	}

	token := loginRes{
		Token: utils.GenerateToken(user.UserId),
		Role:  db.IsAdminRole(user.UserId),
	}
	res.Data = &token
	return res
}

func Register(req *db.TUser) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}

	if ok := utils.RegexpStr(1, req.Name); !ok {
		res.Code = common.FormatError
		res.Msg = "invalid name"
		return res
	}

	req.Passwd = utils.MD5Bytes([]byte(req.Passwd))

	err := db.InsertUser(req)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = errors.WithMessage(err, "insert user failed").Error()
		return res
	}
	res.Data = &req.UserId
	return res
}

func PassReset(req *define.User) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}
	req.Passwd = utils.MD5Bytes([]byte(req.Passwd))
	err := db.UpdatePasswd(req.User, req.Passwd)
	if err != nil {
		res.Code = common.UpdateDBErr
		res.Msg = err.Error()
		return res
	}
	return res
}

func IsValidName(name string) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}
	//utils.RegexpStr(utils.RegexpTypeName, name)
	ok, err := db.IsValidName(name)
	if err != nil {
		res.Code = common.QueryDBErr
		res.Msg = errors.WithMessage(err, "query user name count failed").Error()
		return res
	}
	res.Data = &ok
	return res
}
