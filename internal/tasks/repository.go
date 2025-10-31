package tasks

import "gorm.io/gorm"

type TasksRepository interface {
	GetTasks() ([]Task, error)
	GetTask(id int64) (*Task, error)
	CreateTask(task *Task) error
	UpdateTask(task *Task) error
	DeleteTask(id int64) error
}

type taskRepository struct {
	db *gorm.DB
}

func NewTasksRepository(db *gorm.DB) TasksRepository {
	return &taskRepository{db: db}
}

func (r *taskRepository) GetTasks() ([]Task, error) {
	var tasks []Task
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskRepository) GetTask(id int64) (*Task, error) {
	var task Task
	err := r.db.First(&task, "id = ?", id).Error
	return &task, err
}

func (r *taskRepository) CreateTask(task *Task) error {
	return r.db.Create(&task).Error
}

func (r *taskRepository) UpdateTask(task *Task) error {
	err := r.db.Save(&task).Error
	return err
}

func (r *taskRepository) DeleteTask(id int64) error {
	return r.db.Delete(&Task{}, "id = ?", id).Error
}
