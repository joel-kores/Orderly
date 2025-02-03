package models

type User struct {
	CommonFields
	Name     string  `gorm:"not null"`
	Email    string  `gorm:"unique;not null;unique_index"`
	Phone    string  `gorm:"not null"`
	Password string  `gorm:"not null" json:"password"`
	Orders   []Order `gorm:"foreignKey:UserID"`
}
