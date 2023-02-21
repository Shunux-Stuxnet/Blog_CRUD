package models

type Blog struct {
	Id            int64  `gorm:"primaryKey" json:"id"`
	Title         string `gorm:"type:varchar(300)" json:"title"`
	Description   string `gorm:"type:varchar(1000)" json:"description"`
	Author        string `gorm:"type:varchar(300)" json:"author"`
	PublishedDate string `gorm:"type:date" json:"publishe_date"`
}
