package handler

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rysmaadit/go-template/common/responder"
	"github.com/rysmaadit/go-template/model"
	"gorm.io/gorm"
)

func CreateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var movie model.Movie

		// decode request body ke struct
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		// TODO: insert data movie ke database dan error handling jika slug sudah ada
		if err := db.Create(&movie).Error; err != nil {
			responder.NewHttpResponse(r, w, 422, "slug already exist", nil)
			return
		}

		// return movie yang sudah di insert ke db
		responder.NewHttpResponse(r, w, http.StatusCreated, &movie, nil)
	}
}

func GetMovies(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var movie []model.Movie
		db.Find(&movie)

		// return semua data dari table movie
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)

	}
}

func GetMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var movie model.Movie

		// baca parameter
		params := mux.Vars(r)
		slug := params["slug"]

		/* TODO: ambil data dari db berdasarkan slug dan masukkan di responder dan Check apakah record not found
		 */
		if err := db.Where(&model.Movie{Slug: slug}).First(&movie).Error; err != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, "data not found", nil)
			return
		}

		// return data yang sudah di ambil dari db bedasarkan slug
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}

func UpdateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var movie model.Movie

		// baca parameter
		params := mux.Vars(r)
		slug := params["slug"]

		/* TODO: ambil data dari db berdasarkan slug dan cek apakah record found/not found
		 */
		if err := db.Where(&model.Movie{Slug: slug}).First(&movie).Error; err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, "data not found", nil)
			return
		}

		// decode request body ke struct
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		// error handling unique slug, jika slug sudah ada akan menampilkan error 422
		if err := db.Model(&movie).Where(&model.Movie{Slug: slug}).Updates(&movie).Error; err != nil {
			responder.NewHttpResponse(r, w, 422, "slug already exist", nil)
			return
		}

		// return data yang sudah diupdate
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}

func DeleteMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		var movie model.Movie

		// baca parameter
		params := mux.Vars(r)
		slug := params["slug"]

		/* TODO: ambil data dari db berdasarkan slug dan cek apakah record found / not found
		 */
		if err := db.Where(&model.Movie{Slug: slug}).First(&movie).Error; err != nil {
			responder.NewHttpResponse(r, w, http.StatusNotFound, "data not found", nil)
			return
		}

		// TODO: delete data movie
		db.Where(&model.Movie{Slug: slug}).First(&movie)
		db.Delete(&movie)

		// return movie yang sudah didelete
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}
