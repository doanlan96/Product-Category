package repository

import (
	"errors"
	"fmt"

	"github.com/golang/Product-Category/model"
)

type CategoryRepo struct {
	categories map[int64]*model.Category
	autoID     int64
}

var Categories CategoryRepo

func init() {
	Categories = CategoryRepo{autoID: 0}
	Categories.categories = make(map[int64]*model.Category)
	Categories.InitData("sql:45312")
}

func (r *CategoryRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CategoryRepo) CreateNewCategory(category *model.Category) int64 {
	nextID := r.getAutoID()
	category.Id = nextID
	r.categories[nextID] = category
	return nextID
}

func (r *CategoryRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCategory(&model.Category{
		Name: "Women",
	})
	r.CreateNewCategory(&model.Category{
		Name: "Men",
	})
	r.CreateNewCategory(&model.Category{
		Name: "Kids",
	})
}

func (r *CategoryRepo) GetAllCategories() map[int64]*model.Category {
	return r.categories
}

func (r *CategoryRepo) FindCategoryById(Id int64) (*model.Category, error) {
	if category, ok := r.categories[Id]; ok {
		return category, nil
	} else {
		return nil, errors.New("Category not found")
	}
}
