package repository

import "BelajarGolangUnitTest/entity"

type CategoryRepository interface{
	FindById(id string) *entity.Category
}