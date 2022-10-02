//go:generate mockgen -source ${GOFILE} -destination mocks.go -package $GOPACKAGE
package usecase

import (
	"context"

	"github.com/Moxxx1e/rsoi-2022-lab1-ci-cd-Moxxx1e/internal/models"
)

type repo interface {
	Create(ctx context.Context, person models.Person) (int64, error)
	Update(ctx context.Context, person models.Person) error
	Delete(ctx context.Context, id int64) error
	GetByID(ctx context.Context, id int64) (models.Person, error)
	GetAll(ctx context.Context) (*[]models.Person, error)
}
