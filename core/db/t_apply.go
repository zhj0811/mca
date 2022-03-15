package db

import (
	"fmt"
	"strconv"
	"time"
)

// TApply 申请信息表
type TApply struct {
	ApplyId   int       `json:"apply_id" gorm:"column:apply_id;primaryKey;autoIncrement"`
	WfId      int8      `json:"wf_id"`      //流程id
	WfNodeId  int       `json:"wf_node_id"` //当前待办流程节点id
	PInfoId   int       `json:"p_info_id"`  //个人证书申请信息
	EntOrgId  int       `json:"ent_org_id"`
	EntAntId  int       `json:"ent_ant_id"`
	UserId    string    `json:"user_id"`
	DeviceId  string    `json:"device_id"`
	Status    int8      `json:"status"` //申请状态
	Marked    bool      `json:"-"`      //是否被审核员认领
	CreatedAt time.Time `json:"created_at"`
}

const (
	ApplyStatusFailed  int8 = -1
	ApplyStatusSuccess int8 = 1
)

// GetActApply 根据审核员用户id获取待办事项
func GetActApply(id string) (count int64, res []*TApply, err error) {

	roleSql := fmt.Sprintf("select t_user_role.id from t_user_role where t_user_role.user_id = \"%s\"", id)
	nodeSql := fmt.Sprintf("select t_wf_node_role.wf_node_id from t_wf_node_role where t_wf_node_role.role_id in (%s)", roleSql)

	countSql := fmt.Sprintf("select count(1) from t_apply where t_apply.wf_node_id in (%s)", nodeSql)
	err = db.Raw(countSql).Count(&count).Error
	if err != nil {
		return
	}
	listSql := fmt.Sprintf("select t_apply.* from t_apply where t_apply.wf_node_id in (%s)", nodeSql)

	err = db.Raw(listSql).Scan(&res).Error
	return
}

func InsertApply(row *TApply) error {
	return db.Model(&TApply{}).Create(row).Error
}

func UpdateApplyStatus(id int, status int8) error {
	return db.Model(&TApply{ApplyId: id}).Update("status", status).Error
}

func UpdateApplyNode(id, node int) error {
	return db.Model(&TApply{ApplyId: id}).Update("wf_node_id", node).Error
}

func UpdateApplyMarked(id int, marked bool) error {
	return db.Model(&TApply{ApplyId: id}).Update("marked", marked).Error
}

func GetApplyByApplicant(id string) (res []*TApply, err error) {
	err = db.Model(&TApply{}).Where("user_id = ?", id).Find(&res).Error
	return
}

func MarkApply(id string) error {
	i, err := strconv.Atoi(id)
	if err != nil {
		return err
	}

	err = UpdateApplyMarked(i, true)
	if err != nil {
		return err
	}
	time.Sleep(8 * time.Second)
	return UpdateApplyMarked(i, false)

	//tx := db.Begin()
	//defer func() {
	//	if r := recover(); r != nil {
	//		tx.Rollback()
	//	}
	//}()
	//
	//if err := tx.Error; err != nil {
	//	return err
	//}
	//apply := TApply{}
	////err = tx.Debug().First(&apply).Error
	////if err != nil {
	////	return err
	////}
	//
	//if err := tx.Debug().Clauses(clause.Locking{Strength: "UPDATE"}).First(&apply, i).Error; err != nil {
	//	tx.Rollback()
	//	fmt.Println(err)
	//	return err
	//}
	//fmt.Println("waiting ...")
	//select {}
	//fmt.Println("exit")
	//if err := tx.Commit().Error; err != nil {
	//	tx.Rollback()
	//	return err
	//}
	//return nil
}
