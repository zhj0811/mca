package factory

import (
	"github.com/pkg/errors"
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
)

type CertApplyWfReq struct {
	WfType  int8               `json:"wf_type"`
	WfNodes []CertApplyNodeReq `json:"wf_nodes"`
}

type CertApplyNodeReq struct {
	WfNodeName string `json:"wf_node_name"`
	WfNodeDesc string `json:"wf_node_desc"`
	//WfNodeIndex int8   `json:"wf_node_index"`
	Roles []int8 `json:"roles"`
}

type WfNodeRoleReq struct {
	WfNodeId int    `json:"wf_node_id"`
	Roles    []int8 `json:"roles"`
}

func CreateCertWF(req *CertApplyWfReq) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}
	wf := db.TWorkflow{WftId: req.WfType, WfCount: int8(len(req.WfNodes))}
	err := db.InsertWorkflow(&wf)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = errors.WithMessage(err, "insert workflow failed").Error()
		return res
	}
	for i, node := range req.WfNodes {
		wfNode := db.TWfNode{WfNodeName: node.WfNodeName, WfNodeDesc: node.WfNodeDesc, WfNodeIndex: int8(i), WfId: wf.WfId}
		err = db.InsertWfNode(&wfNode)
		if err != nil {
			err = errors.WithMessage(err, "insert workflow node failed")
			break
		}
		if len(node.Roles) == 0 {
			continue
		}
		var wfNodeRoles []db.TWfNodeRole
		for _, role := range node.Roles {
			wfNodeRole := db.TWfNodeRole{WfNodeId: wfNode.WfNodeId, RoleId: role}
			wfNodeRoles = append(wfNodeRoles, wfNodeRole)
		}
		err = db.InsertWfNodeRoles(&wfNodeRoles)
		if err != nil {
			err = errors.WithMessage(err, "insert workflow node role failed")
			break
		}
	}
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}
	res.Data = &wf
	return res
}

//CreateWfNodeRole 给指定流程节点添加操作组
func CreateWfNodeRole(req *WfNodeRoleReq) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}

	var wfNodeRoles []db.TWfNodeRole
	for _, role := range req.Roles {
		wfNodeRole := db.TWfNodeRole{WfNodeId: req.WfNodeId, RoleId: role}
		wfNodeRoles = append(wfNodeRoles, wfNodeRole)
	}
	err := db.InsertWfNodeRoles(&wfNodeRoles)
	if err != nil {
		res.Msg = errors.WithMessage(err, "insert workflow node role failed").Error()
		res.Code = common.InsertDBErr
		return res
	}
	return res
}
