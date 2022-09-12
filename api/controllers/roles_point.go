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

func (server *Server) CreateRolesPoint(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	rolespoint := models.RolesPoint{}
	err = json.Unmarshal(body, &rolespoint)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	err = rolespoint.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	rolespointCreated, err := rolespoint.SaveRolesPoint(server.DB)

	if err != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, rolespointCreated.ID))
	responses.JSON(w, http.StatusCreated, rolespointCreated)

}

func (server *Server) GetRolesPoints(w http.ResponseWriter, r *http.Request) {

	rolespoint := models.RolesPoint{}

	rolespoints, err := rolespoint.FindAllRolesPoints(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, rolespoints)
}

func (server *Server) GetRolesPoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	rolespoint := models.RolesPoint{}
	rolespointGotten, err := rolespoint.FindRolesPointByID(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, rolespointGotten)
}

func (server *Server) UpdateRolesPoint(w http.ResponseWriter, r *http.Request) {

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
	rolespoint := models.RolesPoint{}
	err = json.Unmarshal(body, &rolespoint)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	err = rolespoint.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedRolesPoint, err := rolespoint.UpdateARolesPoint(server.DB, uint32(kid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedRolesPoint)
}

func (server *Server) DeleteRolesPoint(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	rolespoint := models.RolesPoint{}

	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = rolespoint.DeleteARolesPoint(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", kid))
	responses.JSON(w, http.StatusNoContent, "")
}
