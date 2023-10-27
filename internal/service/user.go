package service

import (
	"context"
	"errors"

	"github.com/as-go/first/internal/models"
)

type service struct {
	store userStorager
}

type Config struct {
	Store userStorager
}

func New(cfg Config) *service {
	s := &service{
		store: cfg.Store,
	}

	return s
}

func (s *service) Create(ctx context.Context, newUser models.CrateUser) (models.User, error) {
	user, _ := s.store.FindByEmail(ctx, newUser.Email)
	if user.ID == 0 {
		return models.User{}, errors.New("email already exist")
	}

	return s.store.Create(ctx, newUser)
}

func (s *service) Delete(ctx context.Context, id int) error {
	user, _ := s.store.FindByID(ctx, id)
	if user.ID == 0 {
		return models.ErrUserNotFound
	}

	return s.store.Delete(ctx, id)
}

func (s *service) FindByID(ctx context.Context, id int) (models.User, error) {
	return s.store.FindByID(ctx, id)
}

func (s *service) Update(ctx context.Context, id int, updUser models.UpdateUser) (models.User, error) {
	user, _ := s.store.FindByID(ctx, id)
	if user.ID == 0 {
		return models.User{}, models.ErrUserNotFound
	}

	return s.store.Update(ctx, id, updUser)
}
