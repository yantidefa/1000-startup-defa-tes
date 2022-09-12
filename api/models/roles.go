package models

import (
	"errors"
	"html"
	"strings"

	"github.com/jinzhu/gorm"
)

type RolesPoint struct {
	ID     int    `json:"id"`
	Roles  string `json:"roles"`
	RPoint int    `json:"r_point"`
}

var ers error

func (s *RolesPoint) Prepare() {
	s.ID = 0
	s.Roles = html.EscapeString(strings.TrimSpace(s.Roles))
}

func (s *RolesPoint) Validate() error {

	if s.Roles == "" {
		return errors.New("required roles")
	}
	if s.RPoint == 0 {
		return errors.New("required point")
	}
	return nil
}

func (s *RolesPoint) SaveRolesPoint(db *gorm.DB) (*RolesPoint, error) {

	ers = db.Debug().Create(&s).Error
	if ers != nil {
		return &RolesPoint{}, ers
	}
	return s, nil
}

func (s *RolesPoint) FindAllRolesPoints(db *gorm.DB) (*[]RolesPoint, error) {
	rolespoints := []RolesPoint{}
	ers = db.Debug().Model(&RolesPoint{}).Limit(100).Find(&rolespoints).Error
	if ers != nil {
		return &[]RolesPoint{}, ers
	}
	return &rolespoints, ers
}

func (s *RolesPoint) FindRolesPointByID(db *gorm.DB, sid uint32) (*RolesPoint, error) {
	ers = db.Debug().Model(RolesPoint{}).Where("id = ?", sid).Take(&s).Error
	if ers != nil {
		return &RolesPoint{}, ers
	}
	if gorm.IsRecordNotFoundError(ers) {
		return &RolesPoint{}, errors.New("kategori buku not found")
	}
	return s, ers
}

func (s *RolesPoint) UpdateARolesPoint(db *gorm.DB, sid uint32) (*RolesPoint, error) {
	db = db.Debug().Model(&RolesPoint{}).Where("id = ?", sid).Take(&RolesPoint{}).UpdateColumns(
		map[string]interface{}{
			"roles":   s.Roles,
			"r_point": s.RPoint,
		},
	)
	if db.Error != nil {
		return &RolesPoint{}, db.Error
	}

	err = db.Debug().Model(&RolesPoint{}).Where("id = ?", sid).Take(&s).Error
	if err != nil {
		return &RolesPoint{}, err
	}
	return s, nil
}

func (s *RolesPoint) DeleteARolesPoint(db *gorm.DB, sid uint32) (int64, error) {

	db = db.Debug().Model(&RolesPoint{}).Where("id = ?", sid).Take(&RolesPoint{}).Delete(&RolesPoint{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
