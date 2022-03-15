package factory

import (
	"strconv"

	"jzsg.com/mca/core/common"
	"jzsg.com/mca/core/db"
)

type HandleSpecApplyReq struct {
	ApplyId  int    `json:"apply_id"`
	WfId     int8   `json:"wf_id"`
	WfNodeId int    `json:"wf_node_id"`
	OprCode  int    `json:"opr_code"`
	Remark   string `json:"remark"`
	UserId   string `json:"user_id"`
}

func HandleSpecApply(req *HandleSpecApplyReq) common.ResponseInfo {
	res := common.ResponseInfo{Code: common.Success}

	opr := db.TWfOperation{
		ApplyId:  req.ApplyId,
		WfNodeId: req.WfNodeId,
		Remark:   req.Remark,
		UserId:   req.UserId,
		OprDesc:  strconv.Itoa(req.OprCode),
	}

	err := db.InsertWfOpr(&opr)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
		return res
	}

	if req.OprCode == -1 {
		err = db.UpdateApplyStatus(req.ApplyId, db.ApplyStatusFailed)
		if err != nil {
			res.Code = common.InsertDBErr
			res.Msg = err.Error()
		}
		return res
	}

	wfNode, err := db.GetWfNodeById(req.WfNodeId)
	if err != nil {
		res.Code = common.QueryDBErr
		res.Msg = err.Error()
		return res
	}
	nextNode, err := db.GetSpecWfNode(wfNode.WfId, wfNode.WfNodeIndex+1)
	if err != nil {
		res.Code = common.QueryDBErr
		res.Msg = err.Error()
		return res
	}

	if nextNode.WfNodeId == 0 {
		//最后一个流程
		err = db.UpdateApplyStatus(req.ApplyId, db.ApplyStatusSuccess)
		if err != nil {
			res.Code = common.UpdateDBErr
			res.Msg = err.Error()
		}
		return res
	}
	err = db.UpdateApplyNode(req.ApplyId, nextNode.WfNodeId)
	if err != nil {
		res.Code = common.InsertDBErr
		res.Msg = err.Error()
	}
	return res
}
