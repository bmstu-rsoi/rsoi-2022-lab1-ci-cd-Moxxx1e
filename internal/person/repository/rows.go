package repository

import "github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"

type dbPersonRow struct {
	ID      int64  `db:"id"`
	Name    string `db:"name"`
	Age     int64  `db:"age"`
	Address string `db:"address"`
	Work    string `db:"work"`
}

func (d *dbPersonRow) toModel() models.Person {
	return models.Person{
		ID:      d.ID,
		Name:    d.Name,
		Age:     d.Age,
		Address: d.Address,
		Work:    d.Work,
	}
}
