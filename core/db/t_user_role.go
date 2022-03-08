package db

type TUserRole struct {
	Id     int    `json:"id" gorm:"column:id;primaryKey; autoIncrement"`
	UserId string `json:"user_id"` //用户名
	RoleId int8   `json:"role_id"` //角色
	Status int8   `json:"status"`  //预留字段
}

func GetAdminId() (id string, err error) {
	err = db.Model(&TUserRole{}).Select("user_id").Where("role_id = ?", RoleAdmin).Limit(1).Find(&id).Error
	return
}

func InsertUserRole(row *TUserRole) error {
	return db.Where(TUserRole{UserId: row.UserId, RoleId: row.RoleId}).FirstOrCreate(row).Error
}
