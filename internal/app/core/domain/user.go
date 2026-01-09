package domain

import (
	"time"
)

// ./internal/app/core/domain/user.go
type User struct {
	ID           int64   `json:"id" gorm:"primaryKey;column:id;autoIncrement"`
	FIO          string  `json:"fio" gorm:"column:fio;type:varchar;not null"`
	RoleID       int64   `json:"roleID" gorm:"column:role_id;not null"`
	Phone        string  `json:"phone" gorm:"column:phone;type:varchar(255);not null;index"`
	ParentID     int64   `json:"parentID" gorm:"column:parent_id;not null"`
	Description  *string `json:"description,omitempty" gorm:"column:description;type:text"`
	PasswordHash string  `json:"password" gorm:"column:password_hash;not null"`
	AuthSecret   string  `json:"-" gorm:"column:auth_secret;not null"`

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;not null;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;not null;autoUpdateTime"`
}

/*
type UserGroup struct {
	ID   int64
	Name string
}
(1,	'Админ'),
(2,	'Менеджер'),
(3,	'Продавец'),
(4,	'Новый');
*/
