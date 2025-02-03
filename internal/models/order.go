package models

type Order struct {
	CommonFields
	UserID     uint        `gorm:"not null"`
	User       User        `gorm:"foreignKey:UserID"`
	TotalPrice float64     `gorm:"not null"`
	OrderItems []OrderItem `gorm:"foreignKey:OrderID"`
}
