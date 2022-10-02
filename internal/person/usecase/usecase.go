package usecase

import (
	"context"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"
)

type UseCase struct {
	repo repo
}

func New(r repo) *UseCase {
	return &UseCase{repo: r}
}

func (u *UseCase) Create(ctx context.Context, person models.Person) (int64, error) {
	return u.repo.Create(ctx, person)
}

func (u *UseCase) Delete(ctx context.Context, id int64) error {
	return u.repo.Delete(ctx, id)
}

func (u *UseCase) Update(ctx context.Context, person models.Person) (models.Person, error) {
	cur, err := u.repo.GetByID(ctx, person.ID)
	if err != nil {
		return models.Person{}, err
	}

	merged := mergePersons(cur, person)
	err = u.repo.Update(ctx, merged)
	if err != nil {
		return models.Person{}, err
	}

	return merged, nil
}

func mergePersons(cur models.Person, update models.Person) models.Person {
	var name, address, work string
	var age int64

	name = cur.Name
	if update.Name != "" {
		name = update.Name
	}
	address = cur.Address
	if update.Address != "" {
		address = update.Address
	}
	work = cur.Work
	if update.Work != "" {
		work = update.Work
	}
	age = cur.Age
	if update.Age != 0 {
		age = update.Age
	}

	return models.Person{
		ID:      cur.ID,
		Name:    name,
		Age:     age,
		Address: address,
		Work:    work,
	}
}

func (u *UseCase) GetByID(ctx context.Context, id int64) (models.Person, error) {
	return u.repo.GetByID(ctx, id)
}

func (u *UseCase) GetAll(ctx context.Context) (*[]models.Person, error) {
	return u.repo.GetAll(ctx)
}
