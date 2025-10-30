package tasks

type TasksService interface {
	GetTasks() ([]Task, error)
	GetTask(id int) (Task, error)
	CreateTask(task Task) (Task, error)
	UpdateTask(id int, task Task) (Task, error)
	DeleteTask(id int) error
}

type taskService struct {
	repo TasksRepository
}

func NewTasksService(repo TasksRepository) TasksService {
	return &taskService{repo: repo}
}

func (s *taskService) GetTasks() ([]Task, error) {
	return s.repo.GetTasks()
}

func (s *taskService) GetTask(id int) (Task, error) {
	return s.repo.GetTask(id)
}

func (s *taskService) CreateTask(task Task) (Task, error) {
	tsk := Task{
		Task:   task.Task,
		Status: task.Status,
	}

	if err := s.repo.CreateTask(tsk); err != nil {
		return Task{}, err
	}

	return tsk, nil
}

func (s *taskService) UpdateTask(id int, task Task) (Task, error) {
	tsk, err := s.repo.GetTask(id)
	if err != nil {
		return Task{}, err
	}

	tsk.Task = task.Task
	tsk.Status = task.Status

	if err := s.repo.UpdateTask(tsk); err != nil {
		return Task{}, err
	}

	return tsk, nil
}

func (s *taskService) DeleteTask(id int) error {
	return s.repo.DeleteTask(id)
}
