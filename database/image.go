package database

import (
	"github.com/hengel2810/api_docli/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"fmt"
	"os"
)

func InsertImage(image models.RequestDockerImage) bool {
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.Dial(mongoURL + ":27017")
	if err != nil {
		fmt.Println(err)
		return false
	}
	defer session.Close()
	c := session.DB("main").C("images")
	err = c.Insert(image)
	if err != nil {
		fmt.Println(err)
		return false
	}
	return true
}

func LoadImages() []models.RequestDockerImage {
	var results []models.RequestDockerImage
	session, err := mgo.Dial("mongo_db:27017")
	if err != nil {
		return results
	}
	defer session.Close()
	c := session.DB("main").C("images")

	err = c.Find(bson.M{}).All(&results)
	if err != nil {
		fmt.Println(err)
		return results
	}
	return results
}
