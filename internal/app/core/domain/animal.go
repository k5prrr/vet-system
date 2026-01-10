package domain

import (
	"time"
)

type Animal struct {
	ID           int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	Name         string    `json:"name" gorm:"column:name;not null"`
	BirthDate    time.Time `json:"birthDate" gorm:"column:birth_date;autoCreateTime"`
	Chip         *string   `json:"chip" gorm:"column:chip;size:255"`
	ClientID     int64     `json:"clientID" gorm:"column:client_id"`
	AnimalTypeID int64     `json:"animalTypeID" gorm:"column:animal_type_id"`
	Description  *string   `json:"description" gorm:"column:description"`
	ParentID     int64     `json:"parentID" gorm:"column:parent_id"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}
