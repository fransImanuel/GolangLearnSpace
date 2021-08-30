package helper

import (
	"fransimanuel/belajargolangrestfulapi/model/domain"
	"fransimanuel/belajargolangrestfulapi/model/web"
)

func ToCategoryResponse(category domain.Category) web.CategoryResponse {
	return web.CategoryResponse{
		Id:   category.Id,
		Name: category.Name,
	}
}

func ToCategoryResponses(categories []domain.Category) []web.CategoryResponse {
	var CategoryResponse []web.CategoryResponse
	for _, category := range categories {
		CategoryResponse = append(CategoryResponse, ToCategoryResponse(category))
	}

	return CategoryResponse

}
