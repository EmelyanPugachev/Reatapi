package taskservice

import "github.com/google/uuid"

type TaskService interface {
	CreateTask(in Taska) (Taska, error)
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
	result := "Call " + task
	return result
}

func (s *taskaService) CreateTask(in Taska) (Taska, error) {
	taska := Taska{
		ID:     uuid.NewString(),
		Task:   s.makeTask(in.Task),
		IsDone: in.IsDone,
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
