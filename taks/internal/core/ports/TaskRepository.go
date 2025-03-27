package ports

import "github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/domain"

type TaskRepository interface {
	FindAll() ([]domain.Task, error)
	//FindByID(id string) (domain.Task, error)
	//Create(task domain.Task) (domain.Task, error)
	//Update(task domain.Task) (domain.Task, error)
	//Delete(id string) error
}

type TaskService interface {
	FindAll() ([]domain.Task, error)
	//FindByID(id string) (domain.Task, error)
	//Create(task domain.Task) (domain.Task, error)
	//Update(task domain.Task) (domain.Task, error)
	//Delete(id string) error
}
