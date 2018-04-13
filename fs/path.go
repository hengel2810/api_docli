package fs

import (
	"os"
	"strings"
)

func TmpDockerImagePath (filename string, userId string) string {
	dir, err := os.Getwd()
	if err != nil {
		return ""
	}
	userDirPath := dir + "/shared/" + userId
	userFolderCreate(userDirPath)
	escapedFileName := strings.Replace(filename, "/", "_", -1)
	fullPath := userDirPath + "/" + escapedFileName
	return fullPath
}

func userFolderCreate(userDirPath string) {
	if _, err := os.Stat(userDirPath); os.IsNotExist(err) {
		os.Mkdir(userDirPath, 0777)
	}
}
