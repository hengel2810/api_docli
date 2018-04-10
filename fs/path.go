package fs

import (
	"os"
	"fmt"
)

func TmpDockerImagePath (filename string) string {
	dir, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	fullPath := dir + "/shared/" + filename
	return fullPath
}
