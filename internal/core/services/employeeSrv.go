package services

import (
	"errors"
	"mansitter/internal/core/domain"
	"mansitter/internal/core/ports"
)

type employeeSrv struct {
	store ports.EmployeeStore
}

/*
	Create([]domain.Employee) error
	Get([]domain.EmployeeID) (error, []domain.Employee)
	Update([]domain.EmployeeOpt) error
	Delete([]domain.EmployeeID) error
*/

func New(eStore ports.EmployeeStore) *employeeSrv {
	return &employeeSrv{
		store: eStore,
	}
}

func validateList([]domain.Employee) error {
	return nil
}

func (srv *employeeSrv) Create(eList []domain.Employee) error {
	if len(eList) < 1 {
		return errors.New("Cannot create employee from an empty list")
	}
	if err := validateList(eList); err != nil {
		return err
	}
	srv.store.Save(eList)
	return nil
}

func (svc *employeeSrv) Get(idList []domain.EmployeeID) (error, []domain.Employee) {
	err, eList := svc.store.Get(idList)
	return err, eList
}

func (svc *employeeSrv) Update(eOptsList []domain.EmployeeOpt) error {
	err := svc.store.Update(eOptsList)
	return err
}

func (svc *employeeSrv) Delete(idList []domain.EmployeeID) error {
	err := svc.store.Delete(idList)
	return err
}
