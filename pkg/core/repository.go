package core

import "context"

type Repository[T interface{}] interface {
	Save(ctx context.Context, entity *T) error
	GetById(ctx context.Context, id string) (*T, error)
}
