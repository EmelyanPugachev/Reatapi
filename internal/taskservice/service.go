package taskservice

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(task string) (Taska, error)
	GetAllTasks() ([]Taska, error)
	GetTaskByID(id string) (Taska, error)
	UpdateTask(id, task string) (Taska, error)
	DeleteTask(id string) error
}

type taskaService struct {
	repo TaskRepository
}

func NewTaskService(r TaskRepository) TaskService {
	return &taskaService{repo: r}
}

func (s *taskaService) makeTask(task string) string {
	result := "Hello " + task
	return result
}

func (s *taskaService) CreateTask(task string) (Taska, error) {
	result := s.makeTask(task)
	taska := Taska{
		ID:   uuid.NewString(),
		Task: result,
	}
	if err := s.repo.CreateTask(taska); err != nil {
		return Taska{}, err
	}
	return taska, nil
}

func (s *taskaService) GetAllTasks() ([]Taska, error) {
	return s.repo.GetAllTasks()
}

func (s *taskaService) GetTaskByID(id string) (Taska, error) {
	return s.repo.GetTaskByID(id)
}

func (s *taskaService) UpdateTask(id, task string) (Taska, error) {
	taska, err := s.repo.GetTaskByID(id)
	if err != nil {
		return Taska{}, err
	}
	result := s.makeTask(task)
	taska.Task = result
	if err := s.repo.UpdateTask(taska); err != nil {
		return Taska{}, err
	}
	return taska, nil
}

func (s *taskaService) DeleteTask(id string) error {
	return s.repo.DeleteTask(id)
}
