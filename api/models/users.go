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

var err error

type User struct {
	ID         int       `json:"id"`
	Name       string    `json:"name"`
	Gender     string    `json:"gender"`
	Birthday   string    `json:"birthday"`
	PhoneNo    string    `json:"phone_no"`
	IDStartup  int       `json:"id_startup"`
	Position   string    `json:"position"`
	IDDomicile int       `json:"id_domicile"`
	Image      string    `json:"image"`
	Email      string    `json:"email"`
	Password   string    `json:"password"`
	CreatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt  time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.Name = html.EscapeString(strings.TrimSpace(u.Name))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Name == "" {
			return errors.New("required name")
		}
		if u.PhoneNo == "" {
			return errors.New("required telp")
		}
		if u.Gender == "" {
			return errors.New("required gender")
		}
		if u.IDStartup == 0 {
			return errors.New("required startup name")
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

	case "login":
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

func (u *User) SaveUser(db *gorm.DB) (*User, *result.Result) {
	const (
		layoutISO = "2006-01-02"
		layoutUS  = "January 2, 2006"
	)
	date := u.Birthday
	fmt.Println(date)
	t, _ := time.Parse(layoutISO, date)
	fmt.Println(t) // 1999-12-31 00:00:00 +0000 UTC
	// fmt.Println(t.Format(layoutUS)) // December 31, 1999
	var user User

	err = db.Debug().Create(&u).Error
	if err != nil {
		res := result.Result{Code: http.StatusInternalServerError, Data: user, Message: err.Error()}
		result, _ := json.Marshal(res)
		fmt.Println(result)
		return &User{}, &res
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {

	users := []User{}
	err = db.Debug().Model(&User{}).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {

	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"name":        u.Name,
			"phone_no":    u.PhoneNo,
			"gender":      u.Gender,
			"birthday":    u.Birthday,
			"id_domicile": u.IDDomicile,
			"id_startup":  u.IDStartup,
			"position":    u.Position,
			"image":       u.Image,
			"password":    u.Password,
			"email":       u.Email,
			"updated_at":  time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
