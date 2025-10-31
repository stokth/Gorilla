package tasks

import (
	"github.com/stretchr/testify/mock"
)

// MockTaskRepository - поддельный репозиторий
type MockTaskRepository struct {
	mock.Mock
}

// DeleteTask implements TasksRepository.
func (m *MockTaskRepository) DeleteTask(id int64) error {
	args := m.Called(id)
	return args.Error(0)
}

// GetTask implements TasksRepository.
// func (m *MockTaskRepository) GetTask(id int64) (*Task, error) {
// 	args := m.Called(id)
// 	var t *Task
// 	if res := args.Get(0); res != nil {
// 		t = res.(*Task)
// 	}
// 	return t, args.Error(1)
// }

func (m *MockTaskRepository) GetTask(id int64) (*Task, error) {
	args := m.Called(id)
	var t *Task
	if res := args.Get(0); res != nil {
		t = res.(*Task)
	}
	return t, args.Error(1)
}

// GetTasks implements TasksRepository.
func (m *MockTaskRepository) GetTasks() ([]Task, error) {
	args := m.Called()
	var tasks []Task
	if res := args.Get(0); res != nil {
		tasks = res.([]Task)
	}
	return tasks, args.Error(1)
}

// UpdateTask implements TasksRepository.
func (m *MockTaskRepository) UpdateTask(task *Task) error {
	args := m.Called(task)
	return args.Error(1)
}

func (m *MockTaskRepository) CreateTask(task *Task) error {
	args := m.Called(task)
	return args.Error(0)
}
