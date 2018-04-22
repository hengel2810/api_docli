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
	"github.com/hengel2810/api_docli/docker"
	"github.com/hengel2810/api_docli/database"
	"github.com/hengel2810/api_docli/helper"
	"fmt"
	"github.com/hengel2810/api_docli/digitalocean"
	"github.com/sirupsen/logrus"
)

func HandlePostImage(w http.ResponseWriter, r *http.Request) {
	decoder := json.NewDecoder(r.Body)
	defer r.Body.Close()
	var docli models.DocliObject
	userId := helper.UserIdFromRequest(r)
	if userId == "" {
		logrus.WithFields(logrus.Fields{"function": "no user id found",}).Error(errors.New("no user id found"))
		handleWriter(w, http.StatusInternalServerError, "no user id found")
		return
	}
	docli.UserId = userId
	err := decoder.Decode(&docli)
	if err == nil {
		validImage := models.DocliObjectValid(docli)
		if validImage {
			docli, err = setDocliData(docli)
			if err == nil {
				id, err := digitalocean.CreateSubdomain(docli.ContainerName)
				if err == nil {
					docli.DomainRecordID = id
					fmt.Println(docli.DomainRecordID)
					err = docker.SetupDocli(docli)
					if err == nil {
						_, err := database.InsertDocli(docli)
						if err == nil {
							w.WriteHeader(http.StatusOK)
						} else {
							logrus.WithFields(logrus.Fields{"function": "database.InsertDocli",}).Error(err)
							handleWriter(w, http.StatusInternalServerError, err.Error())
						}
					} else {
						logrus.WithFields(logrus.Fields{"function": "docker.SetupDocli",}).Error(err)
						handleWriter(w, http.StatusInternalServerError, err.Error())
					}
				} else {
					logrus.WithFields(logrus.Fields{"function": "digitalocean.CreateSubdomain",}).Error(err)
					handleWriter(w, http.StatusInternalServerError, err.Error())
				}
			} else {
				handleWriter(w, http.StatusBadRequest, err.Error())
				logrus.WithFields(logrus.Fields{"function": "setDocliData",}).Error(err)
			}
		} else {
			logrus.WithFields(logrus.Fields{"function": "models.DocliObjectValid",}).Error(errors.New("docli object from request invalid"))
			handleWriter(w, http.StatusBadRequest, err.Error())
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
		logrus.WithFields(logrus.Fields{"function": "HandleGetDoclis",}).Error(errors.New("no user id"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing userId"))
		return
	}
	images, err := database.LoadDoclis(userId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "database.LoadDoclis",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	jsonData, err := json.Marshal(images)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "json.Marshal",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}

func HandleDeleteDocli(w http.ResponseWriter, r *http.Request) {
	docliId := r.URL.Query().Get("docliId")
	if docliId == "" {
		logrus.WithFields(logrus.Fields{"function": "HandleDeleteDocli",}).Error(errors.New("no docli id"))
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte("missing docliId"))
		return
	}
	docli, err := database.DocliFromDocliId(docliId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "database.DocliFromDocliId",}).Error(err)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	err = digitalocean.DeleteSubdomain(docli.DomainRecordID)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "digitalocean.DeleteSubdomain",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = docker.StopContainer(docli)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "docker.StopContainer",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	err = docker.RemoveImage(docli.FullName)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "docker.RemoveImage",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}

	err = database.RemoveDocli(docliId)
	if err != nil {
		logrus.WithFields(logrus.Fields{"function": "database.RemoveDocli",}).Error(err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.WriteHeader(http.StatusOK)
}

