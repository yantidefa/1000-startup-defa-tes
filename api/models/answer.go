package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"net/http"
	"seribu/api/utils/result"
	"strings"

	"github.com/jinzhu/gorm"
)

type Answer struct {
	ID         int    `json:"id"`
	Jawaban    string `json:"answer"`
	IDQuestion string `json:"id_question"`
	Role       int    `json:"id_roles"`
	Point      int    `json:"point"`
}

var ggl error

func (j *Answer) Prepare() {
	j.ID = 0
	j.Jawaban = html.EscapeString(strings.TrimSpace(j.Jawaban))
}

func (j *Answer) Validate() error {

	if j.Jawaban == "" {
		return errors.New("required jawaban")
	}
	if j.Point == 0 {
		return errors.New("required point")
	}
	return nil
}

func (j *Answer) SaveAnswer(db *gorm.DB) (*Answer, *result.Result) {
	var answer Answer
	ggl = db.Debug().Create(&j).Error
	if ggl != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: answer, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Answer{}, &res
	}
	return j, nil
}

func (j *Answer) FindAllAnswers(db *gorm.DB) (*[]Answer, error) {
	Answers := []Answer{}
	ggl = db.Debug().Model(&Answer{}).Limit(100).Find(&Answers).Error
	if ggl != nil {
		return &[]Answer{}, ggl
	}
	return &Answers, ggl
}

func (j *Answer) FindAnswerByID(db *gorm.DB, jid uint32) (*Answer, error) {
	ggl = db.Debug().Model(Answer{}).Where("id = ?", jid).Take(&j).Error
	if ggl != nil {
		return &Answer{}, ggl
	}
	if gorm.IsRecordNotFoundError(ggl) {
		return &Answer{}, errors.New("kategori buku not found")
	}
	return j, ggl
}

func (j *Answer) UpdateAAnswer(db *gorm.DB, jid uint32) (*Answer, error) {
	db = db.Debug().Model(&Answer{}).Where("id = ?", jid).Take(&Answer{}).UpdateColumns(
		map[string]interface{}{
			"jawban":      j.Jawaban,
			"point":       j.Point,
			"id_question": j.IDQuestion,
			"id_roles":    j.Role,
		},
	)
	if db.Error != nil {
		return &Answer{}, db.Error
	}

	err = db.Debug().Model(&Answer{}).Where("id = ?", jid).Take(&j).Error
	if err != nil {
		return &Answer{}, err
	}
	return j, nil
}

func (j *Answer) DeleteAAnswer(db *gorm.DB, jid uint32) (int64, error) {

	db = db.Debug().Model(&Answer{}).Where("id = ?", jid).Take(&Answer{}).Delete(&Answer{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
