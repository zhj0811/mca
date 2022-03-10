package handler

import (
	"github.com/gin-gonic/gin"
	"jzsg.com/mca/common/define"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
	"jzsg.com/mca/core/server/factory"
)

//Login 用户登录
func Login(ctx *gin.Context) {
	req := define.User{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}

	res := factory.Login(req)

	if res.Code != common.Success {
		logger.Errorf("User %s login failed %s.", req.User, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("User %s login success", req.User)
	common.SimpleResponse(ctx, &res)
	return
}

// Register 超管注册操作员
func Register(ctx *gin.Context) {
	req := db.TUser{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.Register(&req)
	if res.Code != common.Success {
		logger.Errorf("Add new operator %s failed %s.", req.Name, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Add new operator %s success", req.Name)
	common.SimpleResponse(ctx, &res)
	return
}

// PassReset 超管重置密码
func PassReset(ctx *gin.Context) {
	req := define.User{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.PassReset(&req)
	if res.Code != common.Success {
		logger.Errorf("Update user %s passwd failed %s.", req.User, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Update user %s passwd success", req.User)
	common.SimpleResponse(ctx, &res)
	return
}

func GetUserInfo(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	res, err := db.GetUserById(id)
	if err != nil {
		logger.Errorf("Get user %s info failed %s.", id, err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Infof("Get user %s info success.", id)
	common.Response(ctx, nil, common.Success, &res)
	return
}

func GetAllOpr(ctx *gin.Context) {
	res, err := db.GetAllOpr()
	if err != nil {
		logger.Errorf("Get operation user info failed %s.", err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Debug("Get operation user info success.")
	common.Response(ctx, nil, common.Success, &res)
	return
}

func PutUserInfo(ctx *gin.Context) {
	req := db.TUser{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	err = db.PutUser(&req)
	if err != nil {
		logger.Errorf("Put user %s info failed %s.", req.UserId, err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Infof("Put user %s info success.", req.UserId)
	common.Response(ctx, nil, common.Success, nil)
	return
}

func IsValidName(ctx *gin.Context) {
	name := ctx.Query("name")
	res := factory.IsValidName(name)
	if res.Code != common.Success {
		//logger.Errorf("Update user %s passwd failed %s.", req.User, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	//logger.Infof("Update user %s passwd success", req.User)
	common.SimpleResponse(ctx, &res)
	return
}
