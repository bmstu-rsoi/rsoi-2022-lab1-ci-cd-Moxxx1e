package usecase

import (
	"context"
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"
	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/person/repository"
)

func TestUseCase_Update(t *testing.T) {
	t.Parallel()
	person := models.Person{
		ID:      1,
		Name:    "Oleg",
		Age:     21,
		Address: "Stroitelnaya 14",
		Work:    "MailRu",
	}
	update := models.Person{
		ID:      1,
		Name:    "",
		Age:     22,
		Address: "Taubenstrasse 3",
		Work:    "Avito",
	}
	merged := models.Person{
		ID:      1,
		Name:    "Oleg",
		Age:     22,
		Address: "Taubenstrasse 3",
		Work:    "Avito",
	}

	tests := []struct {
		name string

		person  models.Person
		prepare func(repo *Mockrepo)

		result models.Person
		err    error
	}{
		{
			name: "positive",

			person: update,
			prepare: func(repo *Mockrepo) {
				repo.EXPECT().GetByID(gomock.Any(), person.ID).Return(person, nil)
				repo.EXPECT().Update(gomock.Any(), merged).Return(nil)
			},

			result: merged,
			err:    nil,
		},
		{
			name: "not found",

			person: update,
			prepare: func(repo *Mockrepo) {
				repo.EXPECT().GetByID(gomock.Any(), person.ID).Return(person, repository.ErrNoPersonWithSuchID)
			},

			result: models.Person{},
			err:    repository.ErrNoPersonWithSuchID,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockrepo(ctrl)
			test.prepare(repo)

			u := New(repo)

			res, err := u.Update(context.Background(), test.person)
			assert.ErrorIs(t, err, test.err)
			assert.Equal(t, test.result, res)
		})
	}
}

func TestUseCase_GetAll(t *testing.T) {
	t.Parallel()
	persons := &[]models.Person{
		{
			ID:      1,
			Name:    "Oleg",
			Age:     22,
			Address: "Stroitelnaya 14",
			Work:    "Avito",
		},
		{
			ID:      1,
			Name:    "Vera",
			Age:     22,
			Address: "Taubenstrasse 3",
			Work:    "Gofore",
		},
	}
	err := errors.New("error from repo")

	tests := []struct {
		name string

		prepare func(repo *Mockrepo)

		result *[]models.Person
		err    error
	}{
		{
			name: "positive",

			prepare: func(repo *Mockrepo) {
				repo.EXPECT().GetAll(gomock.Any()).Return(persons, nil)
			},

			result: persons,
			err:    nil,
		},
		{
			name: "error from repo",

			prepare: func(repo *Mockrepo) {
				repo.EXPECT().GetAll(gomock.Any()).Return(nil, err)
			},

			result: nil,
			err:    err,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			ctrl := gomock.NewController(t)
			defer ctrl.Finish()

			repo := NewMockrepo(ctrl)
			test.prepare(repo)

			u := New(repo)

			res, err := u.GetAll(context.Background())
			assert.ErrorIs(t, err, test.err)
			assert.Equal(t, test.result, res)
		})
	}
}
