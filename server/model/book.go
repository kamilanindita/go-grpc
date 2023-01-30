package model

type BookDB struct {
	Id          int `gorm:"primary_key; auto_increment; not_null"`
	Title       string
	Description string
	Author      string
}

func (e *BookDB) TableName() string {
	return "book"
}
