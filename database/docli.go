package database

import (
	"github.com/hengel2810/api_docli/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"errors"
	"fmt"
)

func InsertDocli(docli models.DocliObject) (bson.ObjectId, error) {
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.Dial(mongoURL + ":27017")
	if err != nil {
		return "", errors.New("error connecting to mongodb")
	}
	defer session.Close()
	collection := session.DB("main").C("images")
	docli.Id = bson.NewObjectId()
	bson.NewObjectId()
	err = collection.Insert(docli)
	if err != nil {
		return "", errors.New("error inserting image to mongodb")
	}
	return docli.Id, nil
}

func RemoveDocli(uniqueId string) error {
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return errors.New("error connecting to mongodb")
	}
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	collection := session.DB("main").C("images")
	fmt.Println("############### PRE REMOVE")
	err = collection.Remove(bson.M{"uniqueid": uniqueId})
	fmt.Println("############### POST REMOVE")
	if err != nil {
		fmt.Println(err)
		return errors.New("remove docli from mongodb fail")
	}
	return nil
}

func LoadDoclis(userId string) ([]models.DocliObject, error) {
	var results []models.DocliObject
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return results, errors.New("error connecting to mongodb")
	}
	defer session.Close()
	c := session.DB("main").C("images")
	err = c.Find(bson.M{"userid": userId}).All(&results)
	if err != nil {
		return results, errors.New("error loading images from mongodb")
	}
	return results, nil
}


func DocliFromDocliId(docliId string) (models.DocliObject, error) {
	var result models.DocliObject
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.Dial(mongoURL)
	if err != nil {
		return result, errors.New("error connecting to mongodb")
	}
	defer session.Close()
	c := session.DB("main").C("images")
	err = c.Find(bson.M{"uniqueid": docliId}).All(&result)
	if err != nil {
		return result, errors.New("error loading images from mongodb")
	}
	return result, nil
}
