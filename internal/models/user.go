package models

type Customer struct {
	CommonFields
	Name  string `gorm:"size:100"`
	Email string `gorm:"uniqueIndex"`
}
