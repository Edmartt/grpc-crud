package tests

import "github.com/edmartt/grpc-test/internal/person/models"

type mockDataAccess struct {
}

func (m mockDataAccess) Create(person models.Person) string {
	return "created"
}

func (m mockDataAccess) Read(id string) (*models.Person, error) {
	return &models.Person{
		ID:        "1",
		FirstName: "ed",
		LastName:  "mart",
		Email:     "email",
	}, nil

}

func (m mockDataAccess) Delete(person *models.Person) (*models.Person, error) {
	return &models.Person{}, nil
}
