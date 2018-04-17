package api

import (
	"net/http"
	"encoding/json"
	"github.com/hengel2810/api_docli/models"
	"github.com/hengel2810/api_docli/database"
	"time"
	"github.com/hengel2810/api_docli/docker"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var docli models.DocliObject
	err := decoder.Decode(&docli)
	if err == nil {
		validImage := models.DocliObjectValid(docli)
		if validImage {
			docli.Uploaded = time.Now()
			dbSuccess := database.InsertImage(docli)
			if dbSuccess == nil {
				err = docker.SetupDocli(docli)
				if err == nil {
					w.WriteHeader(http.StatusOK)
				} else {
					handleWriter(w, http.StatusInternalServerError, "docker setup error")
				}
			} else {
				handleWriter(w, http.StatusInternalServerError, "db error")
			}
		} else {
			handleWriter(w, http.StatusBadRequest, "wrong request object")
		}
	}
}

func handleWriter(w http.ResponseWriter, statusCode int, errorString string) {
	w.WriteHeader(statusCode)
	w.Write([]byte(errorString))
}

