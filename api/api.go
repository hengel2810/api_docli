package api

import (
	"net/http"
	"fmt"
	"encoding/json"
	"github.com/hengel2810/api_docli/models"
	"time"
	"github.com/hengel2810/api_docli/docker"
	"github.com/hengel2810/api_docli/database"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	var image models.RequestDockerImage
	err := decoder.Decode(&image)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
	}
	defer r.Body.Close()
	image.Uploaded = time.Now()
	pulled := docker.PullImage(image.FullName)
	if pulled {
		dbSuccess := database.InsertImage(image)
		if dbSuccess {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusInternalServerError)
	}

}
