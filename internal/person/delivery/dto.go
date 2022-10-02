package delivery

import "github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"

type person struct {
	ID      int64  `json:"id"`
	Name    string `json:"name"`
	Age     int64  `json:"age"`
	Address string `json:"address"`
	Work    string `json:"work"`
}

func fromModel(m models.Person) person {
	return person{
		ID:      m.ID,
		Name:    m.Name,
		Age:     m.Age,
		Address: m.Address,
		Work:    m.Work,
	}
}

func (p person) toModel() models.Person {
	return models.Person{
		ID:      p.ID,
		Name:    p.Name,
		Age:     p.Age,
		Address: p.Address,
		Work:    p.Work,
	}
}

type httpError struct {
	Message string `json:"message"`
}
