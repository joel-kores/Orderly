package models

type Category struct {
	CommonFields
	Name     string     `gorm:"size:100"`
	ParentID *uint      `gorm:"default:null"`
	Parent   *Category  `gorm:"foreignKey:ParentID"`
	Children []Category `gorm:"foreignKey:ParentID"`
}
