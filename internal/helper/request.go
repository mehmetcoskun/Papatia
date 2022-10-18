package helper

import (
	"log"
	"net/http"
)

func Request(url, method string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest(method, url, nil)
	if err != nil {
		log.Fatalln(err)
	}

	req.Header.Set("User-Agent", "Papatia/1.0")

	res, err := client.Do(req)
	if err != nil {
		log.Fatalln(err)
	}

	return res
}
