package tasks

type Task struct {
	ID     uint   `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
}
