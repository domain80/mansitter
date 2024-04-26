package ports

import "mansitter/internal/core/domain"

type EmployeeStore interface {
	Save([]domain.Employee) error
	Get([]domain.EmployeeID) (error, []domain.Employee)
	Update([]domain.EmployeeOpt) error
	Delete([]domain.EmployeeID) error
}

type EmployeeService interface {
	Create([]domain.Employee) error
	Get([]domain.EmployeeID) (error, []domain.Employee)
	Update([]domain.EmployeeOpt) error
	Delete([]domain.EmployeeID) error
}
