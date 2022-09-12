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

func (server *Server) CreateDomicile(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	domicile := models.Domicile{}
	err = json.Unmarshal(body, &domicile)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	domicile.Prepare()
	err = domicile.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	domicileCreated, res := domicile.SaveDomicile(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, domicileCreated.ID))
	responses.JSON(w, http.StatusCreated, domicileCreated)

}

func (server *Server) GetDomiciles(w http.ResponseWriter, r *http.Request) {

	domisili := models.Domicile{}

	domisilis, err := domisili.FindAllDomiciles(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, domisilis)
}

func (server *Server) GetDomicile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	domicile := models.Domicile{}
	domicileGotten, err := domicile.FindDomicileByID(server.DB, uint32(bid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, domicileGotten)
}

func (server *Server) UpdateDomicile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	domicile := models.Domicile{}
	err = json.Unmarshal(body, &domicile)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	domicile.Prepare()
	err = domicile.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedDomicile, err := domicile.UpdateADomicile(server.DB, uint32(bid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedDomicile)
}

func (server *Server) DeleteDomicile(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	blog := models.Blog{}

	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = blog.DeleteABlog(server.DB, uint32(bid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", bid))
	responses.JSON(w, http.StatusNoContent, "")
}
