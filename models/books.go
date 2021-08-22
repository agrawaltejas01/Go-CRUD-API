package models

type Book struct {
	ID     int    `json:"id" gorm:"unique_index"`
	Title  string `json:"title" gorm:"type:varchar(50)"`
	Author string `json:"author" gorm:"type:varchar(50)"`
}
