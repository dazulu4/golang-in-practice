package main

import "testing"

func TestGetFullTimeEmployeeById(t *testing.T) {
	table := []struct {
		id               int
		dni              string
		mockFunc         func()
		expectedEmployee FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{
						Id:       1,
						Position: "CEO",
					}, nil
				}
				GetPersonByDni = func(id string) (Person, error) {
					return Person{
						Name: "John Doe",
						Age:  35,
						DNI:  "1",
					}, nil
				}
			},
			expectedEmployee: FullTimeEmployee{
				Person: Person{
					Age:  35,
					DNI:  "1",
					Name: "John Doe",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonById := GetPersonByDni

	for _, item := range table {
		item.mockFunc()
		ft, err := GetFullTimeEmployeeById(item.id, item.dni)
		if err != nil {
			t.Errorf("Error when getting Employee")
		}
		if ft.Age != item.expectedEmployee.Age {
			t.Errorf("Error, got %d expected %d", ft.Age, item.expectedEmployee.Age)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDni = originalGetPersonById
}
