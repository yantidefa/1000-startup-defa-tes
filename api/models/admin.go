package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"html"
	"log"
	"net/http"
	"seribu/api/utils/result"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// var errs error

type Admin struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Birthday   string    `json:"birthday"`
	PhoneNo    string    `json:"phone_no"`
	IDDomicile int       `json:"id_domicile"`
	Image      string    `json:"image"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hashs(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifiPassword(hasedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hasedPassword), []byte(password))
}

func (u *Admin) BeforeSave() error {
	hasedPassword, err := Hashs(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hasedPassword)
	return nil
}

func (u *Admin) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *Admin) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.PhoneNo == "" {
			return errors.New("required telp")
		}
		if u.IDDomicile == 0 {
			return errors.New("required domisili")
		}
		if u.Image == "" {
			return errors.New("required image")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if u.Birthday == "" {
			return errors.New("required tanggal lahir")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil

	case "logins":
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}
		return nil

	default:
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.Password == "" {
			return errors.New("required password")
		}
		if u.Email == "" {
			return errors.New("required email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("invalid email")
		}

		return nil
	}

}

func (u *Admin) SaveAdmin(db *gorm.DB) (*Admin, *result.Result) {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := u.Birthday
	fmt.Println(date)
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t) // 1999-12-31 00:00:00 +0000 UTC
	// fmt.Println(t.Format(layoutUS)) // December 31, 1999
	var admin Admin

	err = db.Debug().Create(&u).Error
	if err != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: admin, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &Admin{}, &res
	}
	return u, nil
}

func (u *Admin) FindAllAdmins(db *gorm.DB) (*[]Admin, error) {

	admins := []Admin{}
	err = db.Debug().Model(&Admin{}).Find(&admins).Error
	if err != nil {
		return &[]Admin{}, err
	}
	return &admins, err
}

func (u *Admin) FindAdminByID(db *gorm.DB, uid uint32) (*Admin, error) {

	err = db.Debug().Model(Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &Admin{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *Admin) UpdateAAdmin(db *gorm.DB, uid uint32) (*Admin, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).UpdateColumns(
		map[string]interface{}{
			"name":        u.Name,
			"phone_no":    u.PhoneNo,
			"birthday":    u.Birthday,
			"id_domicile": u.IDDomicile,
			"image":       u.Image,
			"password":    u.Password,
			"email":       u.Email,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &Admin{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &Admin{}, err
	}
	return u, nil
}

func (u *Admin) DeleteAAdmin(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&Admin{}).Where("id = ?", uid).Take(&Admin{}).Delete(&Admin{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
