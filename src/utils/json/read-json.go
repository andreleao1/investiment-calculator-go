package readjson

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
)

type FileData struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

func GetValueByKey(key string) (string, error) {
	wd, err := os.Getwd()

	if err != nil {
		return "", errors.New("Error on get current directory")
	}

	file, err := os.Open(filepath.Join(wd, "config", "json", "investments.json"))

	if err != nil {
		return "", errors.New("Error on open file " + "investments.json")
	}

	defer file.Close()

	bytes, err := io.ReadAll(file)

	if err != nil {
		return "", errors.New("Error on read file " + "investments.json")
	}

	var fileData []FileData
	err = json.Unmarshal(bytes, &fileData)

	if err != nil {
		fmt.Println("Aqui ", err)
		return "", errors.New("Error on unmarshal file " + "investments.json")
	}

	for _, data := range fileData {
		if data.Key == key && data.Value != "" {
			return data.Value, nil
		}
	}

	return "", nil
}
