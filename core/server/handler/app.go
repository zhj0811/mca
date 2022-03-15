package handler

import (
	"github.com/gin-gonic/gin"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
	"jzsg.com/mca/core/server/factory"
)

// CreateIndCertApply 创建个人申请
func CreateIndCertApply(ctx *gin.Context) {
	req := factory.IndCertReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.CreateIndCertApply(&req)
	if res.Code != common.Success {
		logger.Errorf("Create individual cert apply %s failed %s.", req.UserId, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Create individual cert apply  %s success", req.UserId)
	common.SimpleResponse(ctx, &res)
	return
}

// CreateEnterCertApply 创建企业证书申请
func CreateEnterCertApply(ctx *gin.Context) {
	req := factory.EnterCertReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.CreateEnterCertApply(&req)
	if res.Code != common.Success {
		logger.Errorf("Create enterprise cert apply %s failed %s.", req.UserId, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Create enterprise cert apply  %s success", req.UserId)
	common.SimpleResponse(ctx, &res)
	return
}

func GetApplicantCert(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	res, err := db.GetApplyByApplicant(id)
	if err != nil {
		logger.Errorf("Get applies by applicant %s failed %s", id, err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	common.Response(ctx, nil, common.Success, &res)
	return
}
