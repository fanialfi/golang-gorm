package model

type User struct {
	ID         uint       `gorm:"primaryKey" json:"id"`
	Name       string     `gorm:"type:varchar(20)" json:"name"`
	Email      string     `gorm:"type:varchar(30)" json:"email"`
	Gender     string     `gorm:"type:varchar(7);default:UNKNOWN" json:"gender"` // jika zero value apapun seperti 0,"",false, akan disimpan sebagai default : UNKNOWN
	Age        uint8      `json:"age"`
	CreditCard CreditCard `gorm:"foreignKey:UserID;references:ID" json:"credit_card"`
}
