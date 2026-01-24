package taskservice

import "gorm.io/gorm"

//

type TaskRepository interface {
	CreateTask(taska Taska) error
	GetAllTasks() ([]Taska, error)
	GetTaskByID(id string) (Taska, error)
	UpdateTask(taska Taska) error
	DeleteTask(id string) error
}

type taskaRepository struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) TaskRepository {
	return &taskaRepository{db: db}
}

func (r *taskaRepository) CreateTask(taska Taska) error {
	return r.db.Create(&taska).Error
}

func (r *taskaRepository) GetAllTasks() ([]Taska, error) {
	var tasks []Taska
	err := r.db.Find(&tasks).Error
	return tasks, err
}

func (r *taskaRepository) GetTaskByID(id string) (Taska, error) {
	var task Taska
	err := r.db.First(&task, "id = ?", id).Error
	return task, err
}

func (r *taskaRepository) UpdateTask(task Taska) error {
	return r.db.Save(&task).Error
}

func (r *taskaRepository) DeleteTask(id string) error {
	return r.db.Delete(&Taska{}, "id = ?", id).Error
}
