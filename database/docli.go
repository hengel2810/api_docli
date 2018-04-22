package database

import (
	"github.com/hengel2810/api_docli/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"errors"
	"time"
)

func InsertDocli(docli models.DocliObject) (bson.ObjectId, error) {
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:[]string{mongoURL + ":27017"},
		Timeout:  60 * time.Second,
		Username: "root",
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return "", errors.New("error connecting to mongodb")
	}
	defer session.Close()
	database := session.DB("main")
	collection := database.C("images")
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
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:[]string{mongoURL + ":27017"},
		Timeout:  60 * time.Second,
		Username: "root",
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return errors.New("error connecting to mongodb")
	}
	defer session.Close()
	session.SetSafe(&mgo.Safe{})
	database := session.DB("main")
	collection := database.C("images")
	err = collection.Remove(bson.M{"uniqueid": uniqueId})
	if err != nil {
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
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:[]string{mongoURL + ":27017"},
		Timeout:  60 * time.Second,
		Username: "root",
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return results, errors.New("error connecting to mongodb")
	}
	defer session.Close()
	database := session.DB("main")
	collection := database.C("images")
	err = collection.Find(bson.M{"userid": userId}).All(&results)
	if err != nil {
		return results, errors.New("error loading images from mongodb")
	}
	return results, nil
}


func DocliFromDocliId(docliId string) (models.DocliObject, error) {
	var result []models.DocliObject
	mongoURL := os.Getenv("MONGOURL")
	if mongoURL == "" {
		mongoURL = "localhost"
	}
	session, err := mgo.DialWithInfo(&mgo.DialInfo{
		Addrs:[]string{mongoURL + ":27017"},
		Timeout:  60 * time.Second,
		Username: "root",
		Password: os.Getenv("MONGO_PASSWORD"),
	})
	if err != nil {
		return models.DocliObject{}, errors.New("error connecting to mongodb")
	}
	defer session.Close()
	database := session.DB("main")
	collection := database.C("images")
	err = collection.Find(bson.M{"uniqueid": docliId}).All(&result)
	if err != nil {
		return models.DocliObject{}, errors.New("error loading images from mongodb")
	}
	if len(result) != 1 {
		return models.DocliObject{}, errors.New("multiple doclis with id in database")
	}
	docliObject := result[0]
	return docliObject, nil
}
