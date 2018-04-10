package fs

import (
	"os"
	"fmt"
	"io"
	"mime/multipart"
)

func CopyImageFromRequest(path string, file multipart.File) error {
	fileHandle, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return err
	}
	_, err = io.Copy(fileHandle, file)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}