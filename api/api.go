package api

import (
	"net/http"
	"github.com/hengel2810/api_docli/docker"
	"fmt"
	"github.com/hengel2810/api_docli/database"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	image, err := DockerImageFromRequest(r)
	if err.StatusCode == 200 {
		err := docker.ImportDockerImage(image)
		if err == nil {
			dbSuccess := database.InsertImage(image)
			if dbSuccess {
				w.WriteHeader(http.StatusOK)
			} else {
				w.WriteHeader(http.StatusInternalServerError)
			}
		} else {
			fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(err.StatusCode)
	}
}
