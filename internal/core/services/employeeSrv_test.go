package services

import (
	"fmt"
	"mansitter/internal/core/domain"
	"testing"
)

func TestCreate(t *testing.T) {
	testEmployees := []domain.Employee{
		{
			ID:         "1",
			Name:       "Another Banger",
			Email:      "clerk@gmail.com",
			Position:   "Accountant",
			Department: "Accounting",
		},
		{
			ID:         "2",
			Name:       "Woman Being",
			Email:      "one@gmail.com",
			Position:   "CEO",
			Department: "Delivery",
		},
	}
	fmt.Println(testEmployees[1])

	// TODO: tesst creating a new emplyee service
	//
	// 	t.Run("create an employee", func(t *testing.T) {
	// 		res := New().Create(testEmployees)
	//
	// 		if res != nil {
	// 			t.Errorf("unable to create an employee")
	// 		}
	// 	})
	//
	// 	t.Run("return an error when an empty array is passed", func(t *testing.T) {
	// 		res := New().Create([]domain.Employee{})
	// 		if res == nil {
	// 			t.Errorf("function should must error out when given an empty array")
	// 		}
	// 	})
}
