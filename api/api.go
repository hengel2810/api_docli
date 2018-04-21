package api

import (
	"net/http"
	"encoding/json"
	"github.com/hengel2810/api_docli/models"
	"time"
	"github.com/Pallinder/sillyname-go"
	"strings"
	"github.com/phayes/freeport"
	"errors"
	"github.com/hengel2810/api_docli/digitalocean"
	"github.com/hengel2810/api_docli/docker"
	"github.com/hengel2810/api_docli/database"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var docli models.DocliObject
	err := decoder.Decode(&docli)
	if err == nil {
		validImage := models.DocliObjectValid(docli)
		if validImage {
			docli, err = setDocliData(docli)
			if err == nil {
				id, err := digitalocean.CreateSubdomain(docli.ContainerName)
				if err == nil {
					docli.DomainRecordID = id
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
				handleWriter(w, http.StatusBadRequest, "error creatin subdomain")
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

func HandleGetDoclis(w http.ResponseWriter, r *http.Request) {
	userId := r.URL.Query().Get("userId")
	if userId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing userId"))
		return
	}
	images, err := database.LoadDoclis(userId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("db error"))
		return
	}
	jsonData, err := json.Marshal(images)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("db error"))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func HandleDeleteDocli(w http.ResponseWriter, r *http.Request) {
	docliId := r.URL.Query().Get("docliId")
	if docliId == "" {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing docliId"))
		return
	}
	docli, err := database.DocliFromDocliId(docliId)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = digitalocean.DeleteSubdomain(docli.ContainerName, docli.DomainRecordID)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = database.RemoveDocli(docliId)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

