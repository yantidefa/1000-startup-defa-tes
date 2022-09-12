package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"seribu/api/utils/result"
	"time"

	"github.com/jinzhu/gorm"
)

type Activity struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Image       string    `json:"image"`
	Location    string    `json:"location"`
	Start       string    `json:"start"`
	Finish      string    `json:"finish"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

var gl error

func (a *Activity) Validate() error {

	if a.Name == "" {
		return errors.New("required name")
	}
	if a.Description == "" {
		return errors.New("required desc")
	}
	if a.Image == "" {
		return errors.New("required img")
	}
	if a.Location == "" {
		return errors.New("required lokasi")
	}
	if a.Start == "" {
		return errors.New("required start")
	}
	if a.Finish == "" {
		return errors.New("required finish")
	}
	return nil
}

func (a *Activity) SaveActivity(db *gorm.DB) (*Activity, *result.Result) {
	const (
		layoutISO = "02-Jan-2006 15:04:05"
		layoutUS  = "00.00"
	)
	datestart := a.Start
	datefinish := a.Finish
	fmt.Println(datestart)
	tstart, _ := time.Parse(layoutISO, datestart)
	tfinish, _ := time.Parse(layoutISO, datefinish)
	fmt.Println(tstart)  // 1999-12-31 00:00:00 +0000 UTC
	fmt.Println(tfinish) // 1999-12-31 00:00:00 +0000 UTC

	var activity Activity
	ersr = db.Debug().Create(&a).Error
	if ersr != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: activity, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Activity{}, &res
	}
	return a, nil
}

func (a *Activity) FindAllActivitys(db *gorm.DB) (*[]Activity, error) {
	Activitys := []Activity{}
	gl = db.Debug().Model(&Activity{}).Limit(100).Find(&Activitys).Error
	if gl != nil {
		return &[]Activity{}, gl
	}
	return &Activitys, gl
}

func (a *Activity) FindActivityByID(db *gorm.DB, aid uint32) (*Activity, error) {
	gl = db.Debug().Model(Activity{}).Where("id = ?", aid).Take(&a).Error
	if gl != nil {
		return &Activity{}, gl
	}
	if gorm.IsRecordNotFoundError(gl) {
		return &Activity{}, errors.New("activity buku not found")
	}
	return a, gl
}

func (a *Activity) UpdateAActivity(db *gorm.DB, aid uint32) (*Activity, error) {
	db = db.Debug().Model(&Activity{}).Where("id = ?", aid).Take(&Activity{}).UpdateColumns(
		map[string]interface{}{
			"name":        a.Name,
			"description": a.Description,
			"image":       a.Image,
			"location":    a.Location,
			"start":       a.Start,
			"finish":      a.Finish,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &Activity{}, db.Error
	}

	err = db.Debug().Model(&Activity{}).Where("id = ?", aid).Take(&a).Error
	if err != nil {
		return &Activity{}, err
	}
	return a, nil
}

func (b *Activity) DeleteAActivity(db *gorm.DB, aid uint32) (int64, error) {

	db = db.Debug().Model(&Activity{}).Where("id = ?", aid).Take(&Activity{}).Delete(&Activity{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
