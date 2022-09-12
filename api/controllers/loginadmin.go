package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"seribu/api/auth"
	"seribu/api/models"
	"seribu/api/responses"
	"seribu/api/utils/formaterror"

	"golang.org/x/crypto/bcrypt"
)

func (server *Server) Logins(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	admin := models.Admin{}
	err = json.Unmarshal(body, &admin)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	admin.Prepare()
	err = admin.Validate("logins")
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	token, err := server.SignIns(admin.Email, admin.Password)
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusUnprocessableEntity, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, token)
}

func (server *Server) SignIns(email, password string) (string, error) {

	var err error

	admin := models.Admin{}

	err = server.DB.Debug().Model(models.Admin{}).Where("email = ?", email).Take(&admin).Error
	if err != nil {
		return "", err
	}
	err = models.VerifiPassword(admin.Password, password)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	return auth.CreateTokenAdmin(uint32(admin.ID))
}
