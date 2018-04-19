package api

import (
	"net/http"
	"encoding/json"
	"github.com/hengel2810/api_docli/models"
	"github.com/hengel2810/api_docli/docker"
	"github.com/hengel2810/api_docli/database"
	"time"
	"github.com/Pallinder/sillyname-go"
	"strings"
	"github.com/phayes/freeport"
	"errors"
	"fmt"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	fmt.Println("############## 333 ################")
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var docli models.DocliObject
	err := decoder.Decode(&docli)
	if err == nil {
		validImage := models.DocliObjectValid(docli)
		if validImage {
			docli, err = setDocliData(docli)
			if err == nil {
				err = docker.SetupDocli(docli)
				if err == nil {
					_, err := database.InsertDocli(docli)
					if err == nil {
						w.WriteHeader(http.StatusOK)
					} else {
						handleWriter(w, http.StatusInternalServerError, err.Error())
					}
				} else {
					handleWriter(w, http.StatusInternalServerError, err.Error())
				}
			} else {
				handleWriter(w, http.StatusInternalServerError, err.Error())
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

func setDocliData(docli models.DocliObject) (models.DocliObject, error) {
	docli.Uploaded = time.Now()
	containerName := sillyname.GenerateStupidName()
	containerName = strings.ToLower(containerName)
	containerName = strings.Replace(containerName, " ", "-", -1)
	containerName = docli.UserId + "-" + containerName
	serverPorts := []models.PortObject{}
	for _, port := range docli.Ports {
		freePort, err := freeport.GetFreePort()
		if err != nil {
			return docli, errors.New("error checking free port")
		}
		serverPort := models.PortObject{Container:port, Host:freePort}
		serverPorts = append(serverPorts, serverPort)
	}
	docli.ServerPorts = serverPorts
	docli.ContainerName = containerName
	return docli, nil
}

