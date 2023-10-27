package service

import (
	"context"

	"github.com/as-go/first/internal/models"
)

type userStorager interface {
	Create(ctx context.Context, user models.CrateUser) (models.User, error)
	Update(ctx context.Context, id int, user models.UpdateUser) (models.User, error)
	FindByID(ctx context.Context, id int) (models.User, error)
	FindByEmail(ctx context.Context, email string) (models.User, error)
	Delete(ctx context.Context, id int) error
}
