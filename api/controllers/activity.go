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

func (server *Server) CreateActivity(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	kegiatan := models.Activity{}
	err = json.Unmarshal(body, &kegiatan)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	// kegiatan.Prepare()
	err = kegiatan.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	kegiatanCreated, res := kegiatan.SaveActivity(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, kegiatanCreated.ID))
	responses.JSON(w, http.StatusCreated, kegiatanCreated)

}

func (server *Server) GetActivitys(w http.ResponseWriter, r *http.Request) {

	kegiatan := models.Activity{}

	kegiatans, err := kegiatan.FindAllActivitys(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, kegiatans)
}

func (server *Server) GetActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	kegiatan := models.Activity{}
	kegiatanGotten, err := kegiatan.FindActivityByID(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, kegiatanGotten)
}

func (server *Server) UpdateActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	kegiatan := models.Activity{}
	err = json.Unmarshal(body, &kegiatan)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	// kegiatan.Prepare()
	err = kegiatan.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedkegiatan, err := kegiatan.UpdateAActivity(server.DB, uint32(pid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedkegiatan)
}

func (server *Server) DeleteActivity(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	kegiatan := models.Activity{}

	pid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = kegiatan.DeleteAActivity(server.DB, uint32(pid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", pid))
	responses.JSON(w, http.StatusNoContent, "")
}
