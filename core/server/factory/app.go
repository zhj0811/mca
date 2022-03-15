package factory

import (
	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
)

type IndCertReq struct {
	db.TPersonCertApplyInfo
	//WfId     int8   `json:"wf_id"` //流程id
	UserId   string `json:"user_id"`
	DeviceId string `json:"device_id"`
}

type EnterCertReq struct {
	db.TEnterApplyAgentInfo
	db.TEnterApplyOrgInfo
	UserId   string `json:"user_id"`
	DeviceId string `json:"device_id"`
}

func CreateIndCertApply(req *IndCertReq) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}
	err := db.InsertIndCertApplyInfo(&req.TPersonCertApplyInfo)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}
	wf, err := db.GetLastWorkflow(db.WfTypePersonalCertVerify)
	if err != nil {
		res.Code = common.QueryDBErr
		res.Msg = err.Error()
		return res
	}
	if wf.WfCount < 1 {
		res.Code = common.InvalidWorkflowErr
		res.Msg = "invalid workflow"
		return res
	}
	node, err := db.GetSpecWfNode(wf.WfId, 1)
	apply := db.TApply{
		WfId:     wf.WfId,
		WfNodeId: node.WfNodeId,
		PInfoId:  req.TPersonCertApplyInfo.PInfoId,
		UserId:   req.UserId,
		DeviceId: req.DeviceId,
	}
	err = db.InsertApply(&apply)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}
	return res
}

func CreateEnterCertApply(req *EnterCertReq) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}
	err := db.InsertAgentInfo(&req.TEnterApplyAgentInfo)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}
	err = db.InsertEnterOrgInfo(&req.TEnterApplyOrgInfo)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}

	wf, err := db.GetLastWorkflow(db.WfTypeEnterpriseCertVerify)
	if err != nil {
		res.Code = common.QueryDBErr
		res.Msg = err.Error()
		return res
	}
	if wf.WfCount < 1 {
		res.Code = common.InvalidWorkflowErr
		res.Msg = "invalid workflow"
		return res
	}
	node, err := db.GetSpecWfNode(wf.WfId, 1)

	apply := db.TApply{
		WfId:     wf.WfId,
		WfNodeId: node.WfNodeId,
		//PInfoId:  req.TPersonCertApplyInfo.PInfoId,
		EntAntId: req.TEnterApplyAgentInfo.AgntId,
		EntOrgId: req.TEnterApplyOrgInfo.EntOrgId,
		UserId:   req.UserId,
		DeviceId: req.DeviceId,
	}
	err = db.InsertApply(&apply)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}
	return res
}
