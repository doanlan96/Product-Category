package repository

import (
	"errors"
	"fmt"

	"github.com/golang/Product-Category/model"
)

type CartRepo struct {
	carts  map[int64]*model.Cart
	autoID int64
}

var Carts CartRepo

func init() {
	Carts = CartRepo{autoID: 0}
	Carts.carts = make(map[int64]*model.Cart)
	Carts.InitData("sql:45312")
}

func (r *CartRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *CartRepo) CreateNewCart(cart *model.Cart) int64 {
	nextID := r.getAutoID()
	cart.Id = nextID
	r.carts[nextID] = cart
	return nextID
}

func (r *CartRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewCart(&model.Cart{
		CategoryID:  2,
		Image:       "/uploads/images/item-02.jpg",
		ProductName: "Herschel supply co 25l",
		Price:       75,
		Quantity:    2,
		CreatedAt:   1614362898000,
		ModifiedAt:  1615410795000})

	r.CreateNewCart(&model.Cart{
		Id:          6,
		CategoryID:  1,
		Image:       "/uploads/images/item-03.jpg",
		ProductName: "Denim jacket blue",
		Price:       92.5,
		Quantity:    1,
		CreatedAt:   1623353412432,
		ModifiedAt:  1623353412432})
}

func (r *CartRepo) GetAllCarts() map[int64]*model.Cart {
	return r.carts
}

func (r *CartRepo) FindCartById(Id int64) (*model.Cart, error) {
	if cart, ok := r.carts[Id]; ok {
		return cart, nil
	} else {
		return nil, errors.New("Cart not found")
	}
}

func (r *CartRepo) DeleteCartById(Id int64) error {
	if _, ok := r.carts[Id]; ok {
		delete(r.carts, Id)
		return nil
	} else {
		return errors.New("Cart not found")
	}
}

func (r *CartRepo) UpdateCartById(newCart *model.Cart) error {
	if _, ok := r.carts[newCart.Id]; ok {
		r.carts[newCart.Id] = newCart
		return nil //tìm được
	} else {
		return errors.New("Cart not found")
	}
}

func (r *CartRepo) UpsertCart(cart *model.Cart) int64 {
	if _, ok := r.carts[cart.Id]; ok {
		r.carts[cart.Id] = cart
		return cart.Id
	} else {
		return r.CreateNewCart(cart)
	}
}
