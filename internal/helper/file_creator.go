package helper

import (
	"encoding/json"
	"os"
)

func CreateFile(fileName string, scrappedData, originalData any) (bool, error) {
	file, err := os.OpenFile("./datas/"+fileName, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return false, err
	}

	data, err := json.MarshalIndent(scrappedData, "", " ")
	if err != nil {
		return false, err
	}

	if _, err := file.Write(data); err != nil {
		return false, err
	}

	if err := file.Close(); err != nil {
		return false, err
	}

	return true, nil
}
