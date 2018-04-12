package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"os"
	"github.com/hengel2810/api_docli/models"
	"github.com/docker/docker/api/types"
)

func ImportDockerImage(requestDockerImage models.RequestDockerImage) error {
	dockerClient, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	reader, err := os.Open(requestDockerImage.Path)
	if err != nil {
		return err
	}
	readCloser, err := dockerClient.ImageImport(context.Background(), types.ImageImportSource{Source:reader, SourceName:"-"}, requestDockerImage.Name,  types.ImageImportOptions{Tag:requestDockerImage.UniqueTag})
	if err != nil {
		return err
	}
	err = readCloser.Close()
	if err != nil {
		return err
	}
	return nil
}