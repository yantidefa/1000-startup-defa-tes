package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type Question struct {
	ID         int    `json:"id"`
	Pertanyaan string `json:"questions"`
}

var er error

func (p *Question) Prepare() {
	p.ID = 0
	p.Pertanyaan = html.EscapeString(strings.TrimSpace(p.Pertanyaan))
}

func (p *Question) Validate() error {

	if p.Pertanyaan == "" {
		return errors.New("required pertanyaan")
	}
	return nil
}

func (p *Question) SaveQuestions(db *gorm.DB) (*Question, error) {

	er = db.Debug().Create(&p).Error
	if er != nil {
		return &Question{}, er
	}
	return p, nil
}

func (p *Question) FindAllQuestions(db *gorm.DB) (*[]Question, error) {
	pertanyaans := []Question{}
	er = db.Debug().Model(&Question{}).Limit(100).Find(&pertanyaans).Error
	if er != nil {
		return &[]Question{}, er
	}
	return &pertanyaans, er
}

func (p *Question) FindQuestionByID(db *gorm.DB, pid uint32) (*Question, error) {
	er = db.Debug().Model(Question{}).Where("id = ?", pid).Take(&p).Error
	if er != nil {
		return &Question{}, er
	}
	if gorm.IsRecordNotFoundError(er) {
		return &Question{}, errors.New("kategori buku not found")
	}
	return p, er
}

func (p *Question) UpdateAQuestion(db *gorm.DB, pid uint32) (*Question, error) {
	db = db.Debug().Model(&Question{}).Where("id = ?", pid).Take(&Question{}).UpdateColumns(
		map[string]interface{}{
			"questions": p.Pertanyaan,
		},
	)
	if db.Error != nil {
		return &Question{}, db.Error
	}

	err = db.Debug().Model(&Question{}).Where("id = ?", pid).Take(&p).Error
	if err != nil {
		return &Question{}, err
	}
	return p, nil
}

func (p *Question) DeleteAQuestion(db *gorm.DB, pid uint32) (int64, error) {

	db = db.Debug().Model(&Question{}).Where("id = ?", pid).Take(&Question{}).Delete(&Question{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
