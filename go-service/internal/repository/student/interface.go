package student

import "go-service/internal/models"

type StudentRepositoryInterface interface {
	GetByID(id string) (*models.StudentDetails, error)
	Insert(s models.StudentDetails) error
	Update(s models.StudentDetails) error
	Delete(id string) error
}
