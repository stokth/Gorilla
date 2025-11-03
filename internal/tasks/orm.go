package tasks

type Task struct {
	ID     int64  `gorm:"primaryKey;autoIncrement" json:"id"`
	Task   string `json:"task"`
	Status string `json:"status"`
	UserID int64  `json:"user_id"`
}

type TaskRequest struct {
	Task   string `json:"task"`
	Status string `json:"status"`
	UserID int64  `json:"user_id"`
}
