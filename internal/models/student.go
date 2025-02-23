package models

type Student struct {
	ID        uint   `gorm:"primaryKey" json:"id"`
	FirstName string `gorm:"type:varchar(32);uniqueIndex:idx_student_fullname" binding:"required,min=3,max=32" json:"first_name"`
	LastName  string `gorm:"type:varchar(32);uniqueIndex:idx_student_fullname" binding:"required,min=3,max=32" json:"last_name"`
	MidName   string `gorm:"type:varchar(32);uniqueIndex:idx_student_fullname" binding:"required,min=3,max=32" json:"mid_name"`
	ClassID   uint   `json:"-"`
	Class     Class  `binding:"required" json:"class"`
}
