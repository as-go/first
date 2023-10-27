package storage

import (
	"context"
	"log"

	"github.com/as-go/first/internal/models"
	"github.com/jackc/pgx/v5/pgxpool"
)

type storage struct {
	pool *pgxpool.Pool
}

type Config struct {
	ConnectionString string
}

func New(cfg Config) *storage {
	p, err := pgxpool.New(context.Background(), cfg.ConnectionString)
	if err != nil {
		log.Fatal(err)
	}

	s := &storage{
		pool: p,
	}

	return s
}

func (s *storage) init() {
	// TODO: создание таблицы
}

func (s *storage) Create(ctx context.Context, user models.CrateUser) (models.User, error) {
	// TODO:
	return models.User{}, nil
}

func (s *storage) Update(ctx context.Context, id int, user models.UpdateUser) (models.User, error) {
	// TODO:
	return models.User{}, nil
}

func (s *storage) FindByID(ctx context.Context, id int) (models.User, error) {
	// TODO:
	return models.User{}, nil
}

func (s *storage) FindByEmail(ctx context.Context, email string) (models.User, error) {
	// TODO:
	return models.User{}, nil
}

func (s *storage) Delete(ctx context.Context, id int) error {
	return nil
}
