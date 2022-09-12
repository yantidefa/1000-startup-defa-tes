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

var EG error

type Books_category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Books_category) Prepare() {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (b *Books_category) Validate() error {

	if b.Name == "" {
		return errors.New("required name")
	}
	if b.Description == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *Books_category) SaveBooks_category(db *gorm.DB) (*Books_category, *result.Result) {
	var books_category Books_category
	ggl = db.Debug().Create(&u).Error
	if ggl != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: books_category, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Books_category{}, &res
	}
	return u, nil
}

func (kb *Books_category) FindAllBooksCategorys(db *gorm.DB) (*[]Books_category, error) {
	kategoribooks := []Books_category{}
	EG = db.Debug().Model(&Books_category{}).Limit(100).Find(&kategoribooks).Error
	if EG != nil {
		return &[]Books_category{}, EG
	}
	return &kategoribooks, EG
}

func (kb *Books_category) FindBooksCategoryByID(db *gorm.DB, kbid uint32) (*Books_category, error) {
	EG = db.Debug().Model(Books_category{}).Where("id = ?", kbid).Take(&kb).Error
	if EG != nil {
		return &Books_category{}, EG
	}
	if gorm.IsRecordNotFoundError(EG) {
		return &Books_category{}, errors.New("kategori buku not found")
	}
	return kb, EG
}

func (kb *Books_category) UpdateABlogsCategory(db *gorm.DB, kbid uint32) (*Books_category, error) {
	db = db.Debug().Model(&Books_category{}).Where("id = ?", kbid).Take(&Books_category{}).UpdateColumns(
		map[string]interface{}{
			"name":        kb.Name,
			"description": kb.Description,
		},
	)
	if db.Error != nil {
		return &Books_category{}, db.Error
	}

	EG = db.Debug().Model(&Books_category{}).Where("id = ?", kbid).Take(&kb).Error
	if EG != nil {
		return &Books_category{}, EG
	}
	return kb, nil
}

func (kb *Books_category) DeleteABooksCategory(db *gorm.DB, kbid uint32) (int64, error) {

	db = db.Debug().Model(&Books_category{}).Where("id = ?", kbid).Take(&Books_category{}).Delete(&Books_category{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
