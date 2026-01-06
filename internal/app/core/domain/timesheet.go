package domain

import "time"

type Timesheet struct {
	ID       int64     `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	UserID   int64     `json:"user_id" gorm:"column:user_id;not null"`
	ParentID int64     `json:"parentID" gorm:"column:parent_id;not null"`
	Date     time.Time `json:"date" gorm:"column:date;not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;not null;autoUpdateTime"`
}
