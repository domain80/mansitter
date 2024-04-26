package domain

type Employee struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Position   string `json:"position"`
	Department string `json:"department"`
}

type EmployeeOpt struct {
		ID         *string `json:"id"`
	Name       *string `json:"name"`
	Email      *string `json:"email"`
	Position   *string `json:"position"`
	Department *string `json:"department"`
}

type EmployeeID string
