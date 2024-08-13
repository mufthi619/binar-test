package repository

import (
	"binar/internal/app/article/category/domain"
)

func ToCategoryEntity(data domain.Category) CategoryEntityGorm {
	return CategoryEntityGorm{
		ID:        data.Id,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func ToCategoriesEntity(data []domain.Category) (finalResponse []CategoryEntityGorm) {
	for _, val := range data {
		finalResponse = append(finalResponse, ToCategoryEntity(val))
	}
	return
}

func ToCategoryDomain(data CategoryEntityGorm) domain.Category {
	return domain.Category{
		Id:        data.ID,
		Name:      data.Name,
		CreatedAt: data.CreatedAt,
		UpdatedAt: data.UpdatedAt,
		DeletedAt: data.DeletedAt,
	}
}

func ToCategoriesDomain(data []CategoryEntityGorm) (finalResponse []domain.Category) {
	for _, val := range data {
		finalResponse = append(finalResponse, ToCategoryDomain(val))
	}
	return
}
