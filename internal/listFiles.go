package internal

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func ListFiles(directory, baseDir string) ([]string, error) {
	var fileList []string

	err := filepath.Walk(directory, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if strings.HasPrefix(path, filepath.Join(directory, "AddOns")) {

			relativePath := strings.TrimPrefix(path, baseDir)

			fileList = append(fileList, relativePath)
		}

		return nil
	})

	fmt.Println(fileList)

	return fileList, err
}
