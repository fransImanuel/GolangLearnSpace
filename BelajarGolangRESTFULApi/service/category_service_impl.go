package service

import (
	"context"
	"database/sql"
	"fransimanuel/belajargolangrestfulapi/execption"
	"fransimanuel/belajargolangrestfulapi/helper"
	"fransimanuel/belajargolangrestfulapi/model/domain"
	"fransimanuel/belajargolangrestfulapi/model/web"
	"fransimanuel/belajargolangrestfulapi/repository"

	"github.com/go-playground/validator/v10"
)

type CategoryServiceImpl struct {
	CategoryRepository repository.CategoryRepository
	DB                 *sql.DB
	validate           *validator.Validate
}

func NewCategoryService(categoryRepository repository.CategoryRepository, db *sql.DB, validate *validator.Validate) CategoryService {
	return &CategoryServiceImpl{
		CategoryRepository: categoryRepository,
		DB:                 db,
		validate:           validate,
	}
}

func (c *CategoryServiceImpl) Create(ctx context.Context, request web.CategoryCreateRequest) web.CategoryResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)

	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category := domain.Category{
		Name: request.Name,
	}

	category = c.CategoryRepository.Save(ctx, tx, category)

	return helper.ToCategoryResponse(category)

}

func (c *CategoryServiceImpl) Update(ctx context.Context, request web.CategoryUpdateRequest) web.CategoryResponse {
	err := c.validate.Struct(request)
	helper.PanicIfError(err)
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, request.Id)
	// helper.PanicIfError(err)
	if err != nil {
		panic(execption.NewNotFoundError(err.Error()))
	}

	category.Name = request.Name

	category = c.CategoryRepository.Update(ctx, tx, category)

	return helper.ToCategoryResponse(category)
}

func (c *CategoryServiceImpl) Delete(ctx context.Context, categoryId int) {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(execption.NewNotFoundError(err.Error()))
	}

	c.CategoryRepository.Delete(ctx, tx, category)

}

func (c *CategoryServiceImpl) FindById(ctx context.Context, categoryId int) web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	category, err := c.CategoryRepository.FindById(ctx, tx, categoryId)
	if err != nil {
		panic(execption.NewNotFoundError(err.Error()))
	}

	return helper.ToCategoryResponse(category)

} //

func (c *CategoryServiceImpl) FindAll(ctx context.Context) []web.CategoryResponse {
	tx, err := c.DB.Begin()
	helper.PanicIfError(err)
	defer helper.CommitOrRollback(tx)

	categories := c.CategoryRepository.FindAll(ctx, tx)

	return helper.ToCategoryResponses(categories)
}
