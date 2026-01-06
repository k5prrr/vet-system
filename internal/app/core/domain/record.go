package domain

import "time"

// Записи на приём
type Record struct {
	ID          int64     `json:"id" gorm:"primaryKey;autoIncrement"`
	TimesheetID int64     `json:"timesheetID" gorm:"column:timesheet_id;not null"`            // Ид расписания
	DateTime    time.Time `json:"dateTime" gorm:"column:date_time;not null;type:timestamptz"` // на какое время

	ClientID int64 `json:"clientID" gorm:"column:client_id;not null"` // клиент
	UserID   int64 `json:"userID" gorm:"column:user_id;not null"`     // На какого врача запись

	ParentID     int64 `json:"parentID" gorm:"column:parent_id;not null"`          // создатель
	ParentRoleID int64 `json:"parentRoleID" gorm:"column:parent_role_id;not null"` // Роль создателя

	StatusID int64 `json:"statusID" gorm:"column:status_id;not null"` // Статус

	AnimalID int64 `json:"animalID" gorm:"column:animal_id;not null"` // Животное

	Complaints      *string `json:"complaints" gorm:"column:complaints;type:text"`                    // Жалобы
	Examination     string  `json:"examination" gorm:"column:examination;not null;type:text"`         // Осмотр/проверка
	DS              *string `json:"ds" gorm:"column:ds;type:text"`                                    // Диагноз
	Recommendations string  `json:"recommendations" gorm:"column:recommendations;not null;type:text"` // Рекомендации
	Description     *string `json:"description" gorm:"column:description;type:text"`                  // Доп описание

	CreatedAt time.Time `json:"createdAt" gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `json:"updatedAt" gorm:"column:updated_at;autoUpdateTime"`
}
