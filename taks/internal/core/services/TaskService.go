package services

import (
	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/domain"
	"github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/ports"
)

type TaskService struct {
	taskRepository ports.TaskRepository
}

func NewTaskService(taskRepository ports.TaskRepository) *TaskService {
	return &TaskService{
		taskRepository: taskRepository,
	}
}

func (s *TaskService) FindAll() ([]domain.Task, error) {
	return s.taskRepository.FindAll()
}
