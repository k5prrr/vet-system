package domain

import (
	"time"
)

type Client struct {
	ID          int64      `json:"id" gorm:"primaryKey;autoIncrement"`
	FIO         string     `json:"fio" gorm:"column:fio;type:varchar;not null"`
	Phone       string     `json:"phone" gorm:"column:phone;type:varchar(255);not null;index"`
	BirthDate   *time.Time `json:"birthDate" gorm:"column:birth_date;autoCreateTime"`
	ParentID    int64      `json:"parentID" gorm:"column:parent_id"`
	Description string     `json:"description" gorm:"column:description;type:text;not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}
