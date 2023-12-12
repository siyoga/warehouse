package ratingdata

import (
	"errors"
	e "warehouseai/ai/errors"
	m "warehouseai/ai/model"

	"github.com/jackc/pgx/v5/pgconn"
	"gorm.io/gorm"
)

type Database struct {
	DB *gorm.DB
}

func (d *Database) Add(rate m.RatingPerUser) *e.DBError {
	if err := d.DB.Create(rate).Error; err != nil {
		if isDuplicateKeyError(err) {
			return e.NewDBError(e.DbExist, "AI with this key/keys already exists.", err.Error())
		} else {
			return e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}
	}

	return nil
}

func (d *Database) Get(conditions map[string]interface{}) (*m.RatingPerUser, *e.DBError) {
	var rate m.RatingPerUser

	if err := d.DB.Where(conditions).First(&rate).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}

		return nil, e.NewDBError(e.DbNotFound, "AI not found.", err.Error())
	}

	return &rate, nil
}

func (d *Database) GetAiRating(id string)

func (d *Database) Update(existRate *m.RatingPerUser, newRate int16) *e.DBError {
	if err := d.DB.Model(existRate).Updates(map[string]interface{}{"rate": newRate}).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}

		return e.NewDBError(e.DbNotFound, "Rate not found.", err.Error())
	}

	return nil
}

func isDuplicateKeyError(err error) bool {
	pgErr, ok := err.(*pgconn.PgError)
	if ok {
		// unique_violation = 23505
		return pgErr.Code == "23505"

	}
	return false
}