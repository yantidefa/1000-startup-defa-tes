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

func (server *Server) CreateBlogsCategory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	blogs_category := models.Blogs_Category{}
	err = json.Unmarshal(body, &blogs_category)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	blogs_category.Prepare()
	err = blogs_category.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	blogs_categoryCreated, res := blogs_category.SaveBlogs_Category(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, blogs_categoryCreated.ID))
	responses.JSON(w, http.StatusCreated, blogs_categoryCreated)

}

func (server *Server) GetBlogsCategorys(w http.ResponseWriter, r *http.Request) {

	blogs_Category := models.Blogs_Category{}

	blogs_categorys, err := blogs_Category.FindAllBlogsCategorys(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, blogs_categorys)
}

func (server *Server) GetBlogsCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	blogs_category := models.Blogs_Category{}
	blog_categoryGotten, err := blogs_category.FindBlogCategoryByID(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, blog_categoryGotten)
}

func (server *Server) UpdateBlogsCategory(w http.ResponseWriter, r *http.Request) {

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
	kategoriblog := models.Blogs_Category{}
	err = json.Unmarshal(body, &kategoriblog)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	kategoriblog.Prepare()
	err = kategoriblog.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedKategoriBlog, err := kategoriblog.UpdateABlogsCategory(server.DB, uint32(kid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedKategoriBlog)
}

func (server *Server) DeleteBlogsCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	kategoriblog := models.Blogs_Category{}

	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = kategoriblog.DeleteAKategoriBlog(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", kid))
	responses.JSON(w, http.StatusNoContent, "")
}
