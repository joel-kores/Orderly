package models

type Product struct {
	CommonFields
	Name       string   `gorm:"not null"`
	Price      float64  `gorm:"not null"`
	CategoryID uint     `gorm:"not null"`
	Category   Category `gorm:"foreignKey:CategoryID"`
}
