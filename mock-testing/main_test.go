package main

import "testing"

func TestGetFillTimeEmployeeById(t *testing.T) {
	table := []struct {
		id       int
		dni      string
		mockFunc func()
		expected FullTimeEmployee
	}{
		{
			id:  1,
			dni: "1",
			mockFunc: func() {
				GetEmployeeById = func(id int) (Employee, error) {
					return Employee{Id: 1, Position: "CEO"}, nil
				}
				GetPersonByDNI = func(dni string) (Person, error) {
					return Person{DNI: "1", Name: "John", Age: 38}, nil
				}
			},
			expected: FullTimeEmployee{
				Person: Person{
					Age:  38,
					DNI:  "1",
					Name: "John",
				},
				Employee: Employee{
					Id:       1,
					Position: "CEO",
				},
			},
		},
	}

	originalGetEmployeeById := GetEmployeeById
	originalGetPersonByDNI := GetPersonByDNI

	for _, test := range table {
		test.mockFunc()
		fte, err := GetFullTimeEmployeeById(test.id, test.dni)

		if err != nil {
			t.Errorf("Error when getting Employee")
		}

		if fte.Age != test.expected.Age {
			t.Errorf("Error, got %d expected %d", fte.Age, test.expected.Age)
		}

		if fte.Name != test.expected.Name {
			t.Errorf("Error, got %v expected %v", fte.Name, test.expected.Name)
		}

		if fte.Position != test.expected.Position {
			t.Errorf("Error, got %v expected %v", fte.Position, test.expected.Position)
		}
	}

	GetEmployeeById = originalGetEmployeeById
	GetPersonByDNI = originalGetPersonByDNI

}
