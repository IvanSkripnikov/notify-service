package models

type Notification struct {
	ID          int    `gorm:"index;type:int" json:"id"`
	UserID      int    `gorm:"index;type:int" json:"userId"`
	Title       string `gorm:"type:text" json:"title"`
	Description string `gorm:"type:text" json:"description"`
	Created     int    `gorm:"index;type:bigint" json:"created"`
}

func (s Notification) TableName() string { return "notifications" }
