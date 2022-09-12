package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"seribu/api/utils/result"
	"strings"
	"time"

	"github.com/jinzhu/gorm"
)

type Blog struct {
	ID              int       `json:"id"`
	Judul           string    `json:"tittle"`
	Konten          string    `json:"conten"`
	Image           string    `json:"image"`
	Id_kategoriblog int       `json:"id_blog_category"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

var ersr error

func (b *Blog) Prepare() {
	b.ID = 0
	b.Judul = html.EscapeString(strings.TrimSpace(b.Judul))
	b.Konten = html.EscapeString(strings.TrimSpace(b.Konten))
	b.Image = html.EscapeString(strings.TrimSpace(b.Image))
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Blog) Validate() error {

	if b.Judul == "" {
		return errors.New("required name")
	}
	if b.Konten == "" {
		return errors.New("required name")
	}
	if b.Image == "" {
		return errors.New("required name")
	}
	if b.Id_kategoriblog == 0 {
		return errors.New("required name")
	}
	return nil
}

func (b *Blog) SaveBlog(db *gorm.DB) (*Blog, *result.Result) {
	var blog Blog
	ersr = db.Debug().Create(&b).Error
	if ersr != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: blog, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Blog{}, &res
	}
	return b, nil
}

func (b *Blog) FindAllBlogs(db *gorm.DB) (*[]Blog, error) {
	Blogs := []Blog{}
	ersr = db.Debug().Model(&Blog{}).Limit(100).Find(&Blogs).Error
	if ersr != nil {
		return &[]Blog{}, ersr
	}
	return &Blogs, ersr
}

func (b *Blog) FindBlogByID(db *gorm.DB, bid uint32) (*Blog, error) {
	ersr = db.Debug().Model(Blog{}).Where("id = ?", bid).Take(&b).Error
	if ersr != nil {
		return &Blog{}, ersr
	}
	if gorm.IsRecordNotFoundError(ersr) {
		return &Blog{}, errors.New("kategori buku not found")
	}
	return b, ersr
}

func (b *Blog) UpdateABlog(db *gorm.DB, bid uint32) (*Blog, error) {
	db = db.Debug().Model(&Blog{}).Where("id = ?", bid).Take(&Blog{}).UpdateColumns(
		map[string]interface{}{
			"judul":            b.Judul,
			"konten":           b.Konten,
			"image":            b.Image,
			"id_blog_category": b.Id_kategoriblog,
			"updated_at":       time.Now(),
		},
	)
	if db.Error != nil {
		return &Blog{}, db.Error
	}

	err = db.Debug().Model(&Blog{}).Where("id = ?", bid).Take(&b).Error
	if err != nil {
		return &Blog{}, err
	}
	return b, nil
}

func (b *Blog) DeleteABlog(db *gorm.DB, bid uint32) (int64, error) {

	db = db.Debug().Model(&Blog{}).Where("id = ?", bid).Take(&Blog{}).Delete(&Blog{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
