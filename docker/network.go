package docker

import (
	"github.com/docker/docker/client"
	"github.com/docker/docker/api/types"
	"golang.org/x/net/context"
	"errors"
)

func HandleNetworks(networks []string, userId string) error {
	for _, networkName := range networks {
		exist, err := networkExists(networkName)
		if err != nil {
			return errors.New("error network exist check")
		} else {
			if exist == false {
				err = createNetwork(networkName)
				if err != nil {
					return errors.New("error create network" + networkName)
				}
			}
		}
	}
	return nil
}

func networkExists(network_name string) (bool, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return  false, err
	}
	networks, err := cli.NetworkList(context.Background(), types.NetworkListOptions{})
	if err != nil {
		return  false, err
	}
	for _, network := range networks {
		if network.Name == network_name {
			return true, nil
		}
	}
	return false, nil
}

func createNetwork(network_name string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		return err
	}
	_, err = cli.NetworkCreate(context.Background(), network_name, types.NetworkCreate{})
	if err != nil {
		return err
	}
	return nil
}
