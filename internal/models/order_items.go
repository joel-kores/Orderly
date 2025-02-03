package models

type OrderItem struct {
	CommonFields
	OrderID   uint    `gorm:"not null"`
	ProductID uint    `gorm:"not null"`
	Quantity  int     `gorm:"not null"`
	Order     Order   `gorm:"foreignKey:OrderID"`
	Product   Product `gorm:"foreignKey:ProductID"`
}
