package domain

type InitData struct {
	User           *User                  `json:"user"`
	UserRoles      map[int64]UserRole     `json:"userRoles"`
	AnimalTypes    map[int64]AnimalType   `json:"animalTypes"`
	RecordStatuses map[int64]RecordStatus `json:"recordStatuses"`
}
