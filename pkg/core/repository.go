package core

import "context"

type Repository[T any] interface {
	Save(ctx context.Context, entity *T) error
	GetById(ctx context.Context, id string) (*T, error)
}
