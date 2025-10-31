package tasks

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     *Task
		mockSetup func(m *MockTaskRepository, input *Task)
		wantErr   bool
	}{
		{
			name:  "успешное создание задачи",
			input: &Task{Task: "Test", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при создании",
			input: &Task{Task: "Bad task", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("CreateTask", input).Return(errors.New("db error"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTaskRepository{}
			tt.mockSetup(mockRepo, tt.input)

			service := NewTasksService(mockRepo)
			result, err := service.CreateTask(tt.input)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTasks(t *testing.T) {
	tests := []struct {
		name      string
		mockSetup func(m *MockTaskRepository)
		wantErr   bool
	}{
		{
			name: "успешное получение задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTasks").Return([]Task{
					{ID: 1, Task: "Task 1", Status: "Pending"},
					{ID: 2, Task: "Task 2", Status: "In Progress"},
				}, nil)
			},
			wantErr: false,
		},
		{
			name: "ошибка при получении задач",
			mockSetup: func(m *MockTaskRepository) {
				m.On("GetTasks").Return(nil, errors.New("ошибка базы данных"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTaskRepository{}
			tt.mockSetup(mockRepo)

			service := NewTasksService(mockRepo)
			result, err := service.GetTasks()

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, result, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestGetTask(t *testing.T) {
	tests := []struct {
		name      string
		input     *Task
		mockSetup func(m *MockTaskRepository, input *Task)
		wantErr   bool
	}{
		{
			name:  "успешно найдена задача",
			input: &Task{ID: 3, Task: "Test", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("GetTask", input.ID).Return(input, nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка: задача не найдена",
			input: &Task{ID: 3, Task: "Bad Test", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("GetTask", input.ID).Return(nil, errors.New("ошибка при обновлении задачи"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTaskRepository{}
			tt.mockSetup(mockRepo, tt.input)

			service := NewTasksService(mockRepo)
			result, err := service.GetTask(tt.input.ID)

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.input, result)
			}

			mockRepo.AssertExpectations(t)
		})
	}
}

func TestUpdateTask(t *testing.T) {
	tests := []struct {
		name      string
		input     *Task
		mockSetup func(m *MockTaskRepository, input *Task)
		wantErr   bool
	}{
		{
			name:  "успешное обновление задачи",
			input: &Task{ID: 3, Task: "Test", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("UpdateTask", input).Return(m.On("GetTask", input.ID).Return(input, nil), nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при обновлении задачи",
			input: &Task{ID: 4, Task: "BadTest", Status: "pending"},
			mockSetup: func(m *MockTaskRepository, input *Task) {
				m.On("GetTask", input.ID).Return(nil, errors.New("ошибка при обновлении задачи"))
			},
			wantErr: true,
		},
	}

	t.Run(tests[0].name, func(t *testing.T) {
		mockRepo := &MockTaskRepository{}
		tests[0].mockSetup(mockRepo, tests[0].input)

		service := NewTasksService(mockRepo)
		result, err := service.UpdateTask(int64(tests[0].input.ID), tests[0].input)

		if tests[0].wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tests[0].input, result)
		}

		mockRepo.AssertExpectations(t)
	})

	t.Run(tests[1].name, func(t *testing.T) {
		mockRepo := &MockTaskRepository{}
		tests[1].mockSetup(mockRepo, tests[1].input)

		service := NewTasksService(mockRepo)
		result, err := service.UpdateTask(int64(tests[1].input.ID), tests[1].input)

		if tests[1].wantErr {
			assert.Error(t, err)
		} else {
			assert.NoError(t, err)
			assert.Equal(t, tests[1].input, result)
		}

		mockRepo.AssertExpectations(t)
	})

	// for _, tt := range tests {

	// }
}

func TestDeleteTask(t *testing.T) {
	tests := []struct {
		name      string
		input     int64
		mockSetup func(m *MockTaskRepository, input int64)
		wantErr   bool
	}{
		{
			name:  "успешное удаление задачи",
			input: 3,
			mockSetup: func(m *MockTaskRepository, input int64) {
				m.On("DeleteTask", input).Return(nil)
			},
			wantErr: false,
		},
		{
			name:  "ошибка при удалении задачи",
			input: 3,
			mockSetup: func(m *MockTaskRepository, input int64) {
				m.On("DeleteTask", input).Return(errors.New("ошибка удаления"))
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := &MockTaskRepository{}
			tt.mockSetup(mockRepo, tt.input)

			service := NewTasksService(mockRepo)
			err := service.DeleteTask(int64(tt.input))

			if tt.wantErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)

			}

			mockRepo.AssertExpectations(t)
		})
	}
}
