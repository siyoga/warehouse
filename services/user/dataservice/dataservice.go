package dataservice

import (
	e "warehouseai/user/errors"
	m "warehouseai/user/model"
)

type UserInterface interface {
	Create(user *m.User) *e.DBError
	RawUpdate(userId string, updatedFields interface{}) (*m.User, *e.DBError) // TODO: Надо избавиться от этого апдейта и перейти на просто Update (siyoga)
	Update(userId string, newValues map[string]interface{}) *e.DBError
	GetOneByPreload(conditions map[string]interface{}, preload string) (*m.User, *e.DBError)
	GetOneBy(conditions map[string]interface{}) (*m.User, *e.DBError)
	Delete(condition map[string]interface{}) *e.DBError
}

type FavoritesInterface interface {
	Add(favorite *m.UserFavorite) *e.DBError
	GetUserFavorites(userId string) (*[]m.UserFavorite, *e.DBError)
	GetFavorite(userId string, aiId string) (*m.UserFavorite, *e.DBError)
	Delete(userId string, aiId string) *e.DBError
}
