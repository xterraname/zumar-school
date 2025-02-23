package models

type Class struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Degree uint8  `gorm:"type:SMALLINT" binding:"required,gte=1,lte=11" json:"degree"`
	Group  string `gorm:"type:varchar(1)" binding:"required" json:"group"`
}
