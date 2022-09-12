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

type Book struct {
	ID              int       `json:"id"`
	Tittle          string    `json:"tittle"`
	Description     string    `json:"description"`
	Buku            string    `json:"book"`
	Edisi           string    `json:"edition"`
	Tahun           string    `json:"publication_year"`
	Author          string    `json:"author"`
	Image           string    `json:"image"`
	Id_kategoribook int       `json:"id_books_category"`
	CreatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt       time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

var ggll error

func (b *Book) Prepare() {
	b.ID = 0
	b.Tittle = html.EscapeString(strings.TrimSpace(b.Tittle))
	b.Description = html.EscapeString(strings.TrimSpace(b.Description))
	b.Image = html.EscapeString(strings.TrimSpace(b.Image))
	b.CreatedAt = time.Now()
	b.UpdatedAt = time.Now()
}

func (b *Book) Validate() error {

	if b.Tittle == "" {
		return errors.New("required name")
	}
	if b.Description == "" {
		return errors.New("required name")
	}
	if b.Image == "" {
		return errors.New("required name")
	}
	if b.Id_kategoribook == 0 {
		return errors.New("required name")
	}
	return nil
}

func (b *Book) SaveBook(db *gorm.DB) (*Book, *result.Result) {
	var book Book
	ggll = db.Debug().Create(&b).Error
	if ggll != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: book, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Book{}, &res
	}
	return b, nil
}

func (b *Book) FindAllBooks(db *gorm.DB) (*[]Book, error) {
	Books := []Book{}
	ggll = db.Debug().Model(&Book{}).Limit(100).Find(&Books).Error
	if ggll != nil {
		return &[]Book{}, ggll
	}
	return &Books, ggll
}

func (b *Book) FindBookByID(db *gorm.DB, bid uint32) (*Book, error) {
	ggll = db.Debug().Model(Book{}).Where("id = ?", bid).Take(&b).Error
	if ggll != nil {
		return &Book{}, ggll
	}
	if gorm.IsRecordNotFoundError(ggll) {
		return &Book{}, errors.New("kategori buku not found")
	}
	return b, ggll
}

func (b *Book) UpdateABook(db *gorm.DB, bid uint32) (*Book, error) {
	db = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&Book{}).UpdateColumns(
		map[string]interface{}{
			"tittle":           b.Tittle,
			"description":      b.Description,
			"book":             b.Buku,
			"edition":          b.Edisi,
			"author":           b.Author,
			"image":            b.Image,
			"id_book_category": b.Id_kategoribook,
			"updated_at":       time.Now(),
		},
	)
	if db.Error != nil {
		return &Book{}, db.Error
	}

	err = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&b).Error
	if err != nil {
		return &Book{}, err
	}
	return b, nil
}

func (b *Book) DeleteABook(db *gorm.DB, bid uint32) (int64, error) {

	db = db.Debug().Model(&Book{}).Where("id = ?", bid).Take(&Book{}).Delete(&Book{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
