package entity

type Role struct {
	Name string `gorm:"type:varchar(50);primaryKey"`
	ID   int    `gorm:"type:int;primaryKey"`
}

func (u *Role) TableName() string {
	return "roles"
}
