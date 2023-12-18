package internal

import (
	"fmt"
	"os"
)

func CheckPath () {
  configFile := "./../config.json"

  v, err := readFromFile(configFile)
  if err != nil {
    fmt.Println("Error while reading file")
  }
}

func readFromFile(file string) (string, error) {
  data, err := os.ReadFile(file)
  if err != nil {
    return "", err
  }

  return string(data), nil
}

func writeToFile(file, value string) error {

}
