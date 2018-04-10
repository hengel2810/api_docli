package docker

import (
	"github.com/docker/docker/client"
	"golang.org/x/net/context"
	"fmt"
	"os"
)

func LoadContainer(path string) error {
	cli, err := client.NewEnvClient()
	if err != nil {
		fmt.Println(err)
		return err
	}
	reader, err := os.Open(path)
	if err != nil {
		fmt.Println(err)
		return err
	}
	resp, err := cli.ImageLoad(context.Background(), reader, false)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = resp.Body.Close()
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}