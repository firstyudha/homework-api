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

		// decode request body ke struct
		var movie model.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		// TODO: insert data movie ke database
		db.Create(&movie)

		// return movie yang sudah di insert ke db
		responder.NewHttpResponse(r, w, http.StatusCreated, &movie, nil)
	}
}

func GetMovies(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var movie []model.Movie
		db.Find(&movie)

		// return data yang sudah di ambil dari db
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}

func GetMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// baca parameter
		params := mux.Vars(r)
		slug := params["slug"]

		// TODO: ambil data dari db berdasarkan slug dan masukkan di responder
		var movie model.Movie
		db.Where(&model.Movie{Slug: slug}).First(&movie)

		// return data yang sudah di ambil dari db
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}

func UpdateMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// decode request body ke struct
		var movie model.Movie
		err := json.NewDecoder(r.Body).Decode(&movie)
		if err != nil {
			responder.NewHttpResponse(r, w, http.StatusBadRequest, nil, err)
			return
		}

		params := mux.Vars(r)
		slug := params["slug"]
		// TODO: update data movie ke database
		db.Where(&model.Movie{Slug: slug}).Save(&movie)

		// return movie yang sudah di insert ke db
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}

func DeleteMovie(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		// decode request body ke struct
		var movie model.Movie
		params := mux.Vars(r)
		slug := params["slug"]
		// TODO: update data movie ke database
		db.Where(&model.Movie{Slug: slug}).First(&movie)

		db.Delete(&movie)

		// return movie yang sudah di insert ke db
		responder.NewHttpResponse(r, w, http.StatusOK, &movie, nil)
	}
}
