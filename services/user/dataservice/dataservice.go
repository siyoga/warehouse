package dataservice

import (
	e "warehouseai/user/errors"
	m "warehouseai/user/model"
)

type UserInterface interface {
	Create(user *m.User) *e.DBError
	RawUpdate(userId string, updatedFields interface{}) (*m.User, *e.DBError)
	GetOneByPreload(conditions map[string]interface{}, preload string) (*m.User, *e.DBError)
	GetOneBy(conditions map[string]interface{}) (*m.User, *e.DBError)
}

type FavoritesInterface interface {
	Add(favorite *m.UserFavorites) *e.DBError
	GetUserFavorites(userId string) (*[]m.UserFavorites, *e.DBError)
	GetFavorite(userId string, aiId string) (*m.UserFavorites, *e.DBError)
	Delete(userId string, aiId string) *e.DBError
}