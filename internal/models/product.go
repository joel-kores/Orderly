package models

type Product struct {
	CommonFields
	Name       string   `gorm:"not null" json:"name"`
	Price      float64  `gorm:"not null" json:"price"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"-"`
}

type GetProducts struct {
	CommonFields
	Name       string   `gorm:"not null" json:"name"`
	Price      float64  `gorm:"not null" json:"price"`
	CategoryID uint     `gorm:"not null" json:"category_id"`
	Category   Category `gorm:"foreignKey:CategoryID" json:"category"`
}
