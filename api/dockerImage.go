package api

import (
	"time"
	"net/http"
	"github.com/hengel2810/api_docli/fs"
	"github.com/hengel2810/api_docli/helper"
	"github.com/hengel2810/api_docli/models"
)

func DockerImageFromRequest(r *http.Request) (models.RequestDockerImage, models.RequestDockerImageError) {
	err := r.ParseForm()
	if err != nil {
		return models.RequestDockerImage{}, models.RequestDockerImageError{StatusCode: 500, Msg:"ParseForm error"}
	}
	file, header, err := r.FormFile("file")
	if err != nil {
		return models.RequestDockerImage{}, models.RequestDockerImageError{StatusCode: 400, Msg:"missing docker image file"}
	}
	defer file.Close()
	userId := helper.UserIdFromRequest(r)
	if userId == "" {
		return models.RequestDockerImage{}, models.RequestDockerImageError{StatusCode: 400, Msg:"missing userid"}
	}
	imagePath := fs.TmpDockerImagePath(header.Filename, userId)
	err = fs.CopyImageFromRequest(imagePath, file)
	if err != nil {
		return models.RequestDockerImage{}, models.RequestDockerImageError{StatusCode: 500, Msg:"error copy image tp tmppath"}
	}
	imageName := r.FormValue("image")
	if imageName == "" || len(imageName) == 0 {
		return models.RequestDockerImage{}, models.RequestDockerImageError{StatusCode: 400, Msg:"missing image name"}
	}
	img := models.RequestDockerImage{Name: imageName, Path: imagePath, UserId:userId, Uploaded: time.Now()}
	return img, models.RequestDockerImageError{StatusCode: 200}
}
