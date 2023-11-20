package model

type CreditCard struct {
	ID     int
	Number string `gorm:"type:varchar(10)"`
	UserID uint   `gorm:"index"`
}
