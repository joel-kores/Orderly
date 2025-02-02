package models

type Product struct {
	CommonFields
	Name       string `gorm:"size:255"`
	Price      float64
	CategoryID uint
	Category   Category
}
