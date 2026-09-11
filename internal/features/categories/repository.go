package categories

import "sprig/internal/model"

type CategoryRepository interface {
	AddCategory(category *model.Category) error
	FindCategoryByName(name string) (*model.Category, error)
	FindByID(id uint64) (*model.Category, error)
	FindAll() ([]model.Category, error)
	Save(category *model.Category) error
	DeleteByID(id uint64) error
}
