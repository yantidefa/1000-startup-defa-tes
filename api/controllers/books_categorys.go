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

func (server *Server) CreateBooksCategory(w http.ResponseWriter, r *http.Request) {

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
	}

	books_category := models.Books_category{}
	err = json.Unmarshal(body, &books_category)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}

	books_category.Prepare()
	err = books_category.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	books_categoryCreated, res := books_category.SaveBooks_category(server.DB)

	if res != nil {

		formattedError := formaterror.FormatError(err.Error())

		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	w.Header().Set("Location", fmt.Sprintf("%s%s/%d", r.Host, r.RequestURI, books_categoryCreated.ID))
	responses.JSON(w, http.StatusCreated, books_categoryCreated)

}

func (server *Server) GetBooksCategorys(w http.ResponseWriter, r *http.Request) {

	books_Category := models.Books_category{}

	books_categorys, err := books_Category.FindAllBooksCategorys(server.DB)
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	responses.JSON(w, http.StatusOK, books_categorys)
}

func (server *Server) GetBooksCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)
	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	books_category := models.Books_category{}
	book_categoryGotten, err := books_category.FindBooksCategoryByID(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	responses.JSON(w, http.StatusOK, book_categoryGotten)
}

func (server *Server) UpdateBooksCategory(w http.ResponseWriter, r *http.Request) {

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
	kategoribook := models.Books_category{}
	err = json.Unmarshal(body, &kategoribook)
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	kategoribook.Prepare()
	err = kategoribook.Validate()
	if err != nil {
		responses.ERROR(w, http.StatusUnprocessableEntity, err)
		return
	}
	updatedKategoriBook, err := kategoribook.UpdateABlogsCategory(server.DB, uint32(kid))
	if err != nil {
		formattedError := formaterror.FormatError(err.Error())
		responses.ERROR(w, http.StatusInternalServerError, formattedError)
		return
	}
	responses.JSON(w, http.StatusOK, updatedKategoriBook)
}

func (server *Server) DeleteBooksCategory(w http.ResponseWriter, r *http.Request) {

	vars := mux.Vars(r)

	kategoribook := models.Books_category{}

	kid, err := strconv.ParseUint(vars["id"], 10, 32)
	if err != nil {
		responses.ERROR(w, http.StatusBadRequest, err)
		return
	}
	_, err = kategoribook.DeleteABooksCategory(server.DB, uint32(kid))
	if err != nil {
		responses.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	w.Header().Set("Entity", fmt.Sprintf("%d", kid))
	responses.JSON(w, http.StatusNoContent, "")
}
