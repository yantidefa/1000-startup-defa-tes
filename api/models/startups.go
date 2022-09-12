package models

import (
	"errors"
	"time"

	"github.com/jinzhu/gorm"
)

var GAGAL error

type Startup struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Startup) Prepare() {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (b *Startup) Validate() error {

	if b.Name == "" {
		return errors.New("required name")
	}
	if b.Description == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *Startup) SaveStartup(db *gorm.DB) (*Startup, error) {

	GAGAL = db.Debug().Create(&u).Error
	if GAGAL != nil {
		return &Startup{}, GAGAL
	}
	return u, nil
}

func (u *Startup) FindAllStartups(db *gorm.DB) (*[]Startup, error) {

	startups := []Startup{}
	err = db.Debug().Model(&Startup{}).Find(&startups).Error
	if err != nil {
		return &[]Startup{}, err
	}
	return &startups, err
}

func (u *Startup) FindStartupByID(db *gorm.DB, uid uint32) (*Startup, error) {

	err = db.Debug().Model(Startup{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Startup{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Startup{}, errors.New("Startup Not Found")
	}
	return u, err
}

func (u *Startup) UpdateAStartup(db *gorm.DB, uid uint32) (*Startup, error) {

	db = db.Debug().Model(&Startup{}).Where("id = ?", uid).Take(&Startup{}).UpdateColumns(
		map[string]interface{}{
			"name":        u.Name,
			"description": u.Description,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &Startup{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Startup{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Startup{}, err
	}
	return u, nil
}

func (u *Startup) DeleteStartup(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Startup{}).Where("id = ?", uid).Take(&Startup{}).Delete(&Startup{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
