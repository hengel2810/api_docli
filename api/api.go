package api

import (
	"net/http"
	"github.com/hengel2810/api_docli/docker"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	image, err := DockerImageFromRequest(r)
	if err.StatusCode == 200 {
		err := docker.ImportDockerImage(image)
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(err.StatusCode)
	}
}
