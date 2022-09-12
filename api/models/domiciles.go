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

var ER error

type Domicile struct {
	ID        int       `json:"id"`
	City      string    `json:"city"`
	Province  string    `json:"province"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Domicile) Prepare() {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (b *Domicile) Validate() error {

	if b.City == "" {
		return errors.New("required name")
	}
	if b.Province == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *Domicile) SaveDomicile(db *gorm.DB) (*Domicile, *result.Result) {
	var domisili Domicile
	ersr = db.Debug().Create(&u).Error
	if ersr != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: domisili, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Domicile{}, &res
	}
	return u, nil
}
func (p *Domicile) FindAllDomiciles(db *gorm.DB) (*[]Domicile, error) {
	domisili := []Domicile{}
	er = db.Debug().Model(&Domicile{}).Find(&domisili).Error
	if er != nil {
		return &[]Domicile{}, er
	}
	return &domisili, er
}

func (p *Domicile) FindDomicileByID(db *gorm.DB, pid uint32) (*Domicile, error) {
	er = db.Debug().Model(Domicile{}).Where("id = ?", pid).Take(&p).Error
	if er != nil {
		return &Domicile{}, er
	}
	if gorm.IsRecordNotFoundError(er) {
		return &Domicile{}, errors.New("kategori buku not found")
	}
	return p, er
}

func (p *Domicile) UpdateADomicile(db *gorm.DB, pid uint32) (*Domicile, error) {
	db = db.Debug().Model(&Domicile{}).Where("id = ?", pid).Take(&Domicile{}).UpdateColumns(
		map[string]interface{}{
			"city":       p.City,
			"province":   p.Province,
			"updated_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &Domicile{}, db.Error
	}

	err = db.Debug().Model(&Domicile{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Domicile{}, err
	}
	return p, nil
}

func (p *Domicile) DeleteADomicile(db *gorm.DB, pid uint32) (int64, error) {

	db = db.Debug().Model(&Domicile{}).Where("id = ?", pid).Take(&Domicile{}).Delete(&Domicile{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
