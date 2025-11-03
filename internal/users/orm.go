package users

import (
	"Gorilla/internal/tasks"
	"time"
)

type User struct {
	ID         int64     `gorm:"primaryKey;autoIncrement" json:"id"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	Created_At time.Time `gorm:"autoCreateTime" json:"created_at"`
	Deleted_At time.Time `json:"deleted_at"`
	Updated_At time.Time `json:"updated_at"`
	Tasks      []tasks.Task
}

type UserRequest struct {
	Email    string `json:"task"`
	Password string `json:"status"`
}
