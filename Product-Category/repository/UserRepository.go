package repository

import (
	"errors"
	"fmt"

	"github.com/golang/Product-Category/model"
)

type UserRepo struct {
	users  map[int64]*model.User
	autoID int64
}

var Users UserRepo

func init() {
	Users = UserRepo{autoID: 0}
	Users.users = make(map[int64]*model.User)
	Users.InitData("sql:45312")
}

func (r *UserRepo) getAutoID() int64 {
	r.autoID += 1
	return r.autoID
}
func (r *UserRepo) CreateNewUser(User *model.User) int64 {
	nextID := r.getAutoID()
	User.Id = nextID
	r.users[nextID] = User
	return nextID
}

func (r *UserRepo) InitData(connection string) {
	fmt.Println("Connect to ", connection)

	r.CreateNewUser(&model.User{
		FirstName:  "Administrator",
		LastName:   "",
		UserName:   "admin",
		Email:      "admin@gmail.com",
		Password:   "admin",
		Avatar:     "https://robohash.org/eaquequasincidunt.png?size=50x50&set=set1",
		Gender:     "Genderfluid",
		Phone:      "933-658-1213",
		Birthday:   "1994-03-23",
		Status:     true,
		CreatedAt:  1609483221000,
		ModifiedAt: 1609483221000})

	r.CreateNewUser(&model.User{
		FirstName:  "Client 1",
		LastName:   "",
		UserName:   "client1",
		Email:      "client1@gmail.com",
		Password:   "client",
		Avatar:     "https://robohash.org/accusantiumminimamagni.png?size=50x50&set=set1",
		Gender:     "Male",
		Phone:      "510-449-7332",
		Birthday:   "2002-03-11",
		Status:     false,
		CreatedAt:  1617440961000,
		ModifiedAt: 1618301961000})
}

func (r *UserRepo) GetAllUsers() map[int64]*model.User {
	return r.users
}

func (r *UserRepo) FindUserById(Id int64) (*model.User, error) {
	if user, ok := r.users[Id]; ok {
		return user, nil
	} else {
		return nil, errors.New("User not found")
	}
}

func (r *UserRepo) DeleteUserById(Id int64) error {
	if _, ok := r.users[Id]; ok {
		delete(r.users, Id)
		return nil
	} else {
		return errors.New("User not found")
	}
}

func (r *UserRepo) UpdateUserById(newUser *model.User) error {
	if _, ok := r.users[newUser.Id]; ok {
		r.users[newUser.Id] = newUser
		return nil //tìm được
	} else {
		return errors.New("User not found")
	}
}

func (r *UserRepo) UpsertUser(user *model.User) int64 {
	if _, ok := r.users[user.Id]; ok {
		r.users[user.Id] = user
		return user.Id
	} else {
		return r.CreateNewUser(user)
	}
}
