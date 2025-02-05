package models

type User struct {
	CommonFields
	Name     string  `gorm:"not null" json:"name"`
	Email    string  `gorm:"unique;not null;unique_index" json:"email"`
	Phone    string  `gorm:"not null" json:"phone"`
	Password string  `gorm:"not null" json:"-"`
	Orders   []Order `gorm:"foreignKey:UserID" json:"orders"`
}
