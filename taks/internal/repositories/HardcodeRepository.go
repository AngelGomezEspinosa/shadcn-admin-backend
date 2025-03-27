package repositories

import "github.com/AngelGomezEspinosa/shadcn-admin-backend/internal/core/domain"

type HardcodeRepository struct {
}

func NewHardcodeRepository() *HardcodeRepository {
	return &HardcodeRepository{}
}

func (r *HardcodeRepository) FindAll() ([]domain.Task, error) {
	return []domain.Task{
		{ID: "1", Title: "Task 1", Status: "Done", Label: "Label 1", Priority: "High"},
		{ID: "2", Title: "Task 2", Status: "Pending", Label: "Label 2", Priority: "Medium"},
		{ID: "3", Title: "Task 3", Status: "Done", Label: "Label 3", Priority: "Low"},
	}, nil
}
