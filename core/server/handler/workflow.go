package handler

import (
	"github.com/gin-gonic/gin"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
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

func DeleteWfNodeRole(ctx *gin.Context) {
	var req int
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	err = db.DeleteWfNodeRole(req)
	if err != nil {
		logger.Errorf("Delete workflow node role %d failed %s.", req, err.Error())
		common.Response(ctx, err, common.DeleteDBErr, nil)
		return
	}
	logger.Errorf("Delete workflow node role %d success.", req)
	common.Response(ctx, nil, common.Success, nil)
	return
}

func GetLastWorkflows(ctx *gin.Context) {
	logger.Debug("Enter last workflows func...")
	roles, err := db.GetLastWorkflows()
	if err != nil {
		logger.Errorf("Query last workflows err %s", err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Info("Query last workflows success.")
	common.Response(ctx, nil, common.Success, roles)
	return
}

func GetSpecWorkflow(ctx *gin.Context) {
	id := ctx.Param("id")
	res := factory.GetSpecWorkflow(id)
	if res.Code != common.Success {
		logger.Errorf("Get spec workflow %s role failed %s.", id, res.Msg)
		common.SimpleResponse(ctx, &res)
		return
	}
	logger.Infof("Get spec workflow %s rolesuccess", id)
	common.SimpleResponse(ctx, &res)
	return
}

func GetActWfs(ctx *gin.Context) {
	id := ctx.GetHeader("id")
	res, err := db.GetActApply(id)
	if err != nil {
		logger.Errorf("Query  workflows err %s", err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Info("Query  workflows success.")
	common.Response(ctx, nil, common.Success, &res)
	return
}
