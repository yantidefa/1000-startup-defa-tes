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

var ERS error

type Blogs_Category struct {
	ID          int       `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	CreatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func (u *Blogs_Category) Prepare() {
	u.ID = 0
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (b *Blogs_Category) Validate() error {

	if b.Name == "" {
		return errors.New("required name")
	}
	if b.Description == "" {
		return errors.New("required name")
	}
	return nil
}

func (u *Blogs_Category) SaveBlogs_Category(db *gorm.DB) (*Blogs_Category, *result.Result) {
	var blog_category Blogs_Category
	ersr = db.Debug().Create(&u).Error
	if ersr != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: blog_category, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Blogs_Category{}, &res
	}
	return u, nil
}

func (kb *Blogs_Category) FindAllBlogsCategorys(db *gorm.DB) (*[]Blogs_Category, error) {
	kategoriblogs := []Blogs_Category{}
	ERS = db.Debug().Model(&Blogs_Category{}).Limit(100).Find(&kategoriblogs).Error
	if ERS != nil {
		return &[]Blogs_Category{}, ERS
	}
	return &kategoriblogs, ERS
}

func (kb *Blogs_Category) FindBlogCategoryByID(db *gorm.DB, kbid uint32) (*Blogs_Category, error) {
	ERS = db.Debug().Model(Blogs_Category{}).Where("id = ?", kbid).Take(&kb).Error
	if ERS != nil {
		return &Blogs_Category{}, ERS
	}
	if gorm.IsRecordNotFoundError(ERS) {
		return &Blogs_Category{}, errors.New("kategori buku not found")
	}
	return kb, ERS
}

func (kb *Blogs_Category) UpdateABlogsCategory(db *gorm.DB, kbid uint32) (*Blogs_Category, error) {
	db = db.Debug().Model(&Blogs_Category{}).Where("id = ?", kbid).Take(&Blogs_Category{}).UpdateColumns(
		map[string]interface{}{
			"name":        kb.Name,
			"description": kb.Description,
		},
	)
	if db.Error != nil {
		return &Blogs_Category{}, db.Error
	}

	ERS = db.Debug().Model(&Blogs_Category{}).Where("id = ?", kbid).Take(&kb).Error
	if ERS != nil {
		return &Blogs_Category{}, ERS
	}
	return kb, nil
}

func (kb *Blogs_Category) DeleteAKategoriBlog(db *gorm.DB, kbid uint32) (int64, error) {

	db = db.Debug().Model(&Blogs_Category{}).Where("id = ?", kbid).Take(&Blogs_Category{}).Delete(&Blogs_Category{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
