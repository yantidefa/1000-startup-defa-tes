package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"seribu/api/models"
	"seribu/api/responses"
	"seribu/api/utils/formaterror"
	"strconv"

	"github.com/gorilla/mux"
)

func (server *Server) CreateAnswer(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	jwb := models.Answer{}
	err = json.Unmarshal(body, &jwb)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	jwb.Prepare()
	err = jwb.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	jwbCreated, res := jwb.SaveAnswer(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, jwbCreated.ID))
	responses.JSON(w, http.StatusCreated, jwbCreated)

}

func (server *Server) GetAnswers(w http.ResponseWriter, r *http.Request) {

	jwb := models.Answer{}

	jwbs, err := jwb.FindAllAnswers(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, jwbs)
}

func (server *Server) GetAnswer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	jwb := models.Answer{}
	jwbGotten, err := jwb.FindAnswerByID(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, jwbGotten)
}

func (server *Server) UpdateAnswer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	jwb := models.Answer{}
	err = json.Unmarshal(body, &jwb)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	jwb.Prepare()
	err = jwb.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedJwb, err := jwb.UpdateAAnswer(server.DB, uint32(kid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedJwb)
}

func (server *Server) DeleteAnswer(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	jwb := models.Answer{}

	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = jwb.DeleteAAnswer(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", kid))
	responses.JSON(w, http.StatusNoContent, "")
}
