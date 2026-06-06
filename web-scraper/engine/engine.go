package engine

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"net/http"
)

func Start(url string) {
	Links = nil
	fmt.Println("Engine start sucessfully!")
	fmt.Println("")

	htmlBytes, err := GetSource(url)
	if err != nil {
		fmt.Printf("Failed to check: %s. Err: %s\n", url, err.Error())
		return
	}

	MatchUrls(bytes.NewReader(htmlBytes), url)

	for n := range Links {
		fmt.Printf("Checking link No %d - %s\n", n+1, Links[n].Url)
		Links[n].Scrape()
	}

	Format(Links)
}

func GetSource(url string) ([]byte, error) {
	fmt.Printf("Traying to get: %s\n\n", url)
	res, err := http.Get(url)
	// fmt.Printf("\nStatus: %s\n", res.Status)
	if err != nil {
		return []byte{}, err
	}

	if res.StatusCode < 200 || res.StatusCode >= 300 {
		return []byte{}, errors.New("Request failed")
	}

	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return []byte{}, err
	}

	return body, nil
}
