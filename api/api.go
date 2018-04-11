package api

import (
	"fmt"
	"net/http"
	"github.com/hengel2810/api_docli/docker"
	"github.com/hengel2810/api_docli/fs"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	defer file.Close()
	imagePath := fs.TmpDockerImagePath(header.Filename)
	err = fs.CopyImageFromRequest(imagePath, file)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	imageName := r.FormValue("image")
	if(len(imageName) > 0) {
		err = docker.LoadImage(imagePath)
		if err == nil {
			w.WriteHeader(http.StatusOK)
		} else {
			w.WriteHeader(http.StatusInternalServerError)
		}
	} else {
		w.WriteHeader(http.StatusBadRequest)
	}
}
