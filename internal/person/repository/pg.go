package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/jmoiron/sqlx"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"
)

var (
	ErrNoPersonWithSuchID = errors.New("no person with such ID")
)

const (
	insertQuery     = `INSERT INTO persons(name, age, address, work) VALUES($1, $2, $3, $4) RETURNING id`
	updateQuery     = `UPDATE persons SET name = $1, age = $2, address = $3, work = $4 WHERE id = $5`
	selectByIDQuery = `SELECT *
                       FROM persons
					   WHERE id = $1`
	selectAll = `SELECT *
			     FROM persons`
	deleteQuery = `DELETE FROM persons WHERE id = $1`
)

type PG struct {
	db *sqlx.DB
}

func NewPG(db *sqlx.DB) *PG {
	return &PG{db: db}
}

func (p *PG) Create(ctx context.Context, person models.Person) (int64, error) {
	row := p.db.QueryRowxContext(ctx, insertQuery, person.Name, person.Age, person.Address, person.Work)

	var id int64
	err := row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (p *PG) Delete(ctx context.Context, id int64) error {
	_, err := p.db.ExecContext(ctx, deleteQuery, id)
	return err
}

func (p *PG) GetByID(ctx context.Context, id int64) (models.Person, error) {
	row := p.db.QueryRowxContext(ctx, selectByIDQuery, id)

	var person dbPersonRow
	err := row.StructScan(&person)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Person{}, ErrNoPersonWithSuchID
		}
		return models.Person{}, err
	}

	return person.toModel(), nil
}

func (p *PG) GetAll(ctx context.Context) (*[]models.Person, error) {
	rows, err := p.db.QueryxContext(ctx, selectAll)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.Person
	for rows.Next() {
		var person dbPersonRow
		err = rows.StructScan(&person)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person.toModel())
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}

	return &persons, nil
}

func (p *PG) Update(ctx context.Context, person models.Person) error {
	_, err := p.db.ExecContext(ctx, updateQuery, person.Name, person.Age, person.Address, person.Work, person.ID)
	return err
}
