package docker

import (
	"github.com/hengel2810/api_docli/models"
)

func SetupDocli(docli models.DocliObject) error {
	err := PullImage(docli.FullName)
	if err != nil {
		return err
	} else {
		err = StartContainer(docli)
		if err != nil {
			RemoveImage(docli.FullName)
			return err
		}
	}
	return err
}

func RemoveDocli(docli models.DocliObject) error {
	err := RemoveImage(docli.FullName)
	if err != nil {
		return err
	} else {
		err = RemoveDocli(docli)
	}
	return err
}
