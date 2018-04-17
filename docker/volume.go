package docker

import (
	"errors"
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types/volume"
	"golang.org/x/net/context"
	"github.com/docker/docker/api/types/filters"
)

func HandleVolumes(volumes []string, userId string) error {
	for _, volumeName := range volumes {
		exist, err := volumeExists(volumeName)
		if err != nil {
			return errors.New("error volume exist check")
		} else {
			if exist == false {
				err = createVolume(volumeName)
				if err != nil {
					return errors.New("error create volume " + volumeName)
				}
			}
		}
	}
	return nil
}

func volumeExists(volume_name string) (bool, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return  false, err
	}
	volumes, err := cli.VolumeList(context.Background(), filters.Args{})
	if err != nil {
		return  false, err
	}
	for _, volume := range volumes.Volumes {
		if volume.Name == volume_name {
			return true, nil
		}
	}
	return false, nil
}

func createVolume(volume_name string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	_, err = cli.VolumeCreate(context.Background(), volume.VolumesCreateBody{Driver: "", Name: volume_name})
	if err != nil {
		return err
	}
	return nil
}