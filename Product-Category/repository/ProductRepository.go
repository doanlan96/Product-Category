package repository

import (
	"errors"
	"fmt"

	"github.com/golang/Product-Category/model"
)

type ProductRepo struct {
	products map[int64]*model.Product
	autoID   int64
}

var Products ProductRepo

func init() {
	Products = ProductRepo{autoID: 0}
	Products.products = make(map[int64]*model.Product)
	Products.InitData("sql:45312")
}

func (r *ProductRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *ProductRepo) CreateNewProduct(product *model.Product) int64 {
	nextID := r.getAutoID()
	product.Id = nextID
	r.products[nextID] = product
	return nextID
}

func (r *ProductRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewProduct(&model.Product{
		CategoryID: 2,
		Images: []string{
			"/uploads/images/item-02.jpg",
			"/uploads/images/item-02.jpg",
			"/uploads/images/item-02.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 1, Comment: "Hay", Rating: 4},
		// 	{ProductId: 1, Comment: "Hay", Rating: 5},
		// 	{ProductId: 1, Comment: "Hay", Rating: 3},
		// 	{ProductId: 1, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Herschel supply co 25l",
		Price:         75,
		IsSale:        true,
		CreatedAt:     1614362898000,
		ModifiedAt:    1615410795000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 1,
		Images: []string{
			"/uploads/images/item-03.jpg",
			"/uploads/images/item-03.jpg",
			"/uploads/images/item-03.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 2, Comment: "Hay", Rating: 4},
		// 	{ProductId: 2, Comment: "Hay", Rating: 5},
		// 	{ProductId: 2, Comment: "Hay", Rating: 3},
		// 	{ProductId: 2, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Denim jacket blue",
		Price:         92.5,
		IsSale:        false,
		CreatedAt:     1610281342000,
		ModifiedAt:    1619283693000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 3,
		Images: []string{
			"/uploads/images/item-04.jpg",
			"/uploads/images/item-04.jpg",
			"/uploads/images/item-04.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 3, Comment: "Hay", Rating: 4},
		// 	{ProductId: 3, Comment: "Hay", Rating: 5},
		// 	{ProductId: 3, Comment: "Hay", Rating: 3},
		// 	{ProductId: 3, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Coach slim easton black",
		Price:         165.9,
		IsSale:        false,
		CreatedAt:     1615745962000,
		ModifiedAt:    1615976362000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 1,
		Images: []string{
			"/uploads/images/item-05.jpg",
			"/uploads/images/item-05.jpg",
			"/uploads/images/item-05.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 4, Comment: "Hay", Rating: 4},
		// 	{ProductId: 4, Comment: "Hay", Rating: 5},
		// 	{ProductId: 4, Comment: "Hay", Rating: 3},
		// 	{ProductId: 4, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Frayed denim shorts",
		Price:         15.9,
		IsSale:        true,
		CreatedAt:     1615746962000,
		ModifiedAt:    1615977362000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 2,
		Images: []string{
			"/uploads/images/item-02.jpg",
			"/uploads/images/item-02.jpg",
			"/uploads/images/item-02.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 5, Comment: "Hay", Rating: 4},
		// 	{ProductId: 5, Comment: "Hay", Rating: 5},
		// 	{ProductId: 5, Comment: "Hay", Rating: 3},
		// 	{ProductId: 5, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Herschel supply co 25l",
		Price:         75,
		IsSale:        false,
		CreatedAt:     1614362898000,
		ModifiedAt:    1615410795000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 1,
		Images: []string{
			"/uploads/images/item-03.jpg",
			"/uploads/images/item-03.jpg",
			"/uploads/images/item-03.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 6, Comment: "Hay", Rating: 4},
		// 	{ProductId: 6, Comment: "Hay", Rating: 5},
		// 	{ProductId: 6, Comment: "Hay", Rating: 3},
		// 	{ProductId: 6, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Denim jacket blue",
		Price:         92.5,
		IsSale:        false,
		CreatedAt:     1610281342000,
		ModifiedAt:    1619283693000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 3,
		Images: []string{
			"/uploads/images/item-04.jpg",
			"/uploads/images/item-04.jpg",
			"/uploads/images/item-04.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 7, Comment: "Hay", Rating: 4},
		// 	{ProductId: 7, Comment: "Hay", Rating: 5},
		// 	{ProductId: 7, Comment: "Hay", Rating: 3},
		// 	{ProductId: 7, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Coach slim easton black",
		Price:         165.9,
		IsSale:        false,
		CreatedAt:     1615745962000,
		ModifiedAt:    1615976362000})

	r.CreateNewProduct(&model.Product{
		CategoryID: 1,
		Images: []string{
			"/uploads/images/item-05.jpg",
			"/uploads/images/item-05.jpg",
			"/uploads/images/item-05.jpg",
		},
		// Reviews: []model.Review{
		// 	{ProductId: 8, Comment: "Hay", Rating: 4},
		// 	{ProductId: 8, Comment: "Hay", Rating: 5},
		// 	{ProductId: 8, Comment: "Hay", Rating: 3},
		// 	{ProductId: 8, Comment: "Hay", Rating: 4},
		// },
		AverageRating: 0,
		Name:          "Frayed denim shorts",
		Price:         15.9,
		IsSale:        false,
		CreatedAt:     1615746962000,
		ModifiedAt:    1615977362000})
}

func (r *ProductRepo) GetAllProducts() map[int64]*model.Product {
	return r.products
}

func (r *ProductRepo) FindProductById(Id int64) (*model.Product, error) {
	if product, ok := r.products[Id]; ok {
		return product, nil
	} else {
		return nil, errors.New("Product not found")
	}
}

func (r *ProductRepo) DeleteProductById(Id int64) error {
	if _, ok := r.products[Id]; ok {
		delete(r.products, Id)
		return nil
	} else {
		return errors.New("Product not found")
	}
}

func (r *ProductRepo) UpdateProductById(newProduct *model.Product) error {
	if _, ok := r.products[newProduct.Id]; ok {
		r.products[newProduct.Id] = newProduct
		return nil //tìm được
	} else {
		return errors.New("Product not found")
	}
}

func (r *ProductRepo) UpsertProduct(product *model.Product) int64 {
	if _, ok := r.products[product.Id]; ok {
		r.products[product.Id] = product
		return product.Id
	} else {
		return r.CreateNewProduct(product)
	}
}

func AverageRating(reviews []model.Review) (result float32) {
	var sum int = 0
	var count int = 0
	var averageRating float32
	for _, review := range reviews {
		sum += review.Rating
		count++
		if sum == 0 || count == 0 {
			averageRating = 0
		} else {
			averageRating = float32(sum / count)
		}
		result = averageRating
	}
	return result
}
