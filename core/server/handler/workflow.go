package handler

import (
	"github.com/gin-gonic/gin"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/server/factory"
)

func CreateCertWF(ctx *gin.Context) {
	req := factory.CertApplyWfReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.CreateCertWF(&req)
	if res.Code != common.Success {
		logger.Errorf("Create cert %d apply workflow failed %s.", req.WfType, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Create cert %d apply workflow success", req.WfType)
	common.SimpleResponse(ctx, &res)
	return
}

func CreateWfNodeRole(ctx *gin.Context) {
	req := factory.WfNodeRoleReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	res := factory.CreateWfNodeRole(&req)
	if res.Code != common.Success {
		logger.Errorf("Create workflow node %d role failed %s.", req.WfNodeId, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Create workflow node %d rolesuccess", req.WfNodeId)
	common.SimpleResponse(ctx, &res)
	return
}
