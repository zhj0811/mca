package handler

import (
	"github.com/gin-gonic/gin"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
	"jzsg.com/mca/core/server/factory"
)

// CreateRole 超管创建角色
func CreateRole(ctx *gin.Context) {
	req := db.TRole{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	err = db.InsertRole(&req)
	if err != nil {
		logger.Errorf("Insert role %s err %s", req.RoleName, err.Error())
		common.Response(ctx, err, common.InsertDBErr, nil)
		return
	}
	logger.Infof("Insert role %s success.", req.RoleName)
	common.Response(ctx, nil, common.Success, &req.RoleId)
	return
}

func CreateUserRole(ctx *gin.Context) {
	req := db.TUserRole{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}

	err = db.InsertUserRole(&req)
	if err != nil {
		logger.Errorf("Insert user role info err %s", err.Error())
		common.Response(ctx, err, common.InsertDBErr, nil)
		return
	}
	logger.Infof("Insert user %s role %d success.", req.UserId, req.RoleId)
	common.Response(ctx, nil, common.Success, &req.Id)
	return
}

func AddUsersForRole(ctx *gin.Context) {
	req := factory.UsersForRoleReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	for _, u := range req.User {
		err = db.InsertUserRole(&db.TUserRole{RoleId: req.Role, UserId: u})
		if err != nil {
			break
		}
	}
	if err != nil {
		logger.Errorf("Insert user role info err %s", err.Error())
		common.Response(ctx, err, common.InsertDBErr, nil)
		return
	}
	logger.Infof("Insert uses %v role %d success.", req.User, req.Role)
	common.Response(ctx, nil, common.Success, nil)
	return
}

func AddRolesForUser(ctx *gin.Context) {
	req := factory.RolesForUserReq{}
	err := ctx.ShouldBindJSON(&req)
	if err != nil {
		logger.Errorf("Request info err %s", err.Error())
		common.Response(ctx, err, common.RequestInfoErr, nil)
		return
	}
	for _, role := range req.Role {
		err = db.InsertUserRole(&db.TUserRole{RoleId: role, UserId: req.User})
		if err != nil {
			break
		}
	}
	if err != nil {
		logger.Errorf("Insert user role info err %s", err.Error())
		common.Response(ctx, err, common.InsertDBErr, nil)
		return
	}
	logger.Infof("Insert uses %s roles %v success.", req.User, req.Role)
	common.Response(ctx, nil, common.Success, nil)
	return
}

func GetRoles(ctx *gin.Context) {
	logger.Debug("Enter get roles func...")
	roles, err := db.GetRoles()
	if err != nil {
		logger.Errorf("Query user roles err %s", err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Info("Query user roles success.")
	common.Response(ctx, nil, common.Success, roles)
	return
}

func GetSpecUserRole(ctx *gin.Context) {
	id := ctx.Param("id")
	roles, err := db.GetSpecUserRole(id)
	if err != nil {
		logger.Errorf("Query user %s roles err %s", id, err.Error())
		common.Response(ctx, err, common.QueryDBErr, nil)
		return
	}
	logger.Infof("Query user %s roles success.", id)
	common.Response(ctx, nil, common.Success, &roles)
	return
}
