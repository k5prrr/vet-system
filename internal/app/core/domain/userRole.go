package domain

import "time"

type UserRole struct {
	ID   int64  `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	Name string `json:"name" gorm:"column:name;not null"`
	Code string `json:"code" gorm:"column:code;not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;not null;autoUpdateTime"`
}
