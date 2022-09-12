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

func (server *Server) CreateBlog(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	blog := models.Blog{}
	err = json.Unmarshal(body, &blog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	blog.Prepare()
	err = blog.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	blogCreated, res := blog.SaveBlog(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, blogCreated.ID))
	responses.JSON(w, http.StatusCreated, blogCreated)

}

func (server *Server) GetBlogs(w http.ResponseWriter, r *http.Request) {

	blog := models.Blog{}

	blogs, err := blog.FindAllBlogs(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, blogs)
}

func (server *Server) GetBlog(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	bid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	blog := models.Blog{}
	blogGotten, err := blog.FindBlogByID(server.DB, uint32(bid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, blogGotten)
}

func (server *Server) UpdateBlog(w http.ResponseWriter, r *http.Request) {

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
	blog := models.Blog{}
	err = json.Unmarshal(body, &blog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	blog.Prepare()
	err = blog.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedBlog, err := blog.UpdateABlog(server.DB, uint32(bid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedBlog)
}

func (server *Server) DeleteBlog(w http.ResponseWriter, r *http.Request) {

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
