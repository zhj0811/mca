package db

type TValidIde struct {
	Id   int    `json:"id" gorm:"autoIncrement;primaryKey"`
	Name string `json:"name"`
	Ide  string `json:"ide" gorm:"uniqueIndex"`
}
