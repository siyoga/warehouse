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

func (d *Database) Add(rate *m.AiRate) *e.DBError {
	if err := d.DB.Create(rate).Error; err != nil {
		if isDuplicateKeyError(err) {
			return e.NewDBError(e.DbExist, "AI with this key/keys already exists.", err.Error())
		} else {
			return e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}
	}

	return nil
}

func (d *Database) Get(conditions map[string]interface{}) (*m.AiRate, *e.DBError) {
	var rate m.AiRate

	if err := d.DB.Where(conditions).First(&rate).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}

		return nil, e.NewDBError(e.DbNotFound, "Rating not found.", err.Error())
	}

	return &rate, nil
}

func (d *Database) GetAverageAiRating(aiId string) (*float64, *e.DBError) {
	var result float64

	if err := d.DB.Model(&m.AiRate{}).Select("COALESCE(AVG(rate), 0)").Where("ai_id = ?", aiId).Scan(&result).Error; err != nil {
		return nil, e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
	}

	return &result, nil
}

func (d *Database) GetCountAiRating(aiId string) (*int64, *e.DBError) {
	var result int64

	if err := d.DB.Model(&m.AiRate{}).Where("ai_id = ?", aiId).Distinct("user_id").Count(&result).Error; err != nil {
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, e.NewDBError(e.DbSystem, "Something went wrong.", err.Error())
		}

		return nil, e.NewDBError(e.DbNotFound, "Rating not found.", err.Error())
	}

	return &result, nil
}

func (d *Database) Update(existRate *m.AiRate, newRate int16) *e.DBError {
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
