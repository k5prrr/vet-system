package domain

type Report struct {
	ClientsTotal     int64 `json:"clientsTotal"`
	RecordsTotal     int64 `json:"recordsTotal"`
	RecordsCancelled int64 `json:"recordsCancelled"`
}
