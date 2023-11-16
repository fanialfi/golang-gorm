package model

type User struct {
	ID    uint   `gorm:"primaryKey"`
	Name  string `gorm:"type:varchar(20)"`
	Email string
	Age   uint8
}
