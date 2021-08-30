package repository

import (
	"context"
	"database/sql"
	"fransimanuel/belajargolangrestfulapi/model/domain"
)

type CategoryRepository interface {
	Save(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, Category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, Category domain.Category)
	FindById(ctx context.Context, tx *sql.Tx, CategoryId int) (domain.Category, error)
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
