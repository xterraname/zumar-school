package models

type Class struct {
	ID     uint   `gorm:"primaryKey" json:"id"`
	Degree uint8  `gorm:"type:SMALLINT;uniqueIndex:idx_degree_group" binding:"required,gte=1,lte=11" json:"degree"`
	Group  string `gorm:"type:varchar(1);uniqueIndex:idx_degree_group" binding:"required" json:"group"`
}
