package docker

import (
	"github.com/hengel2810/api_docli/models"
)

func SetupDocli(docli models.DocliObject) error {
	err := StartContainer(docli)
	return err
}
