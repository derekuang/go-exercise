package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

const (
	ApiKey = "e0bc1cb8"
	URL    = "http://www.omdbapi.com/"
)

type MovieSearchResult struct {
	Title    string
	Poster   string
	Response string
	Error    string
}

func main() {
	result, err := searchMovie(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	posterURL := result.Poster
	posterPath := "./posters/"
	postername := result.Title + ".jpg"
	if err = downloadFile(posterURL, posterPath+postername); err != nil {
		log.Fatal(err)
	}
}

func searchMovie(title string) (*MovieSearchResult, error) {
	resp, err := http.Get(URL + "?apikey=" + ApiKey + "&t=" + title)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return nil, fmt.Errorf("search query failed: %s", resp.Status)
	}

	var result MovieSearchResult
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		resp.Body.Close()
		return nil, err
	}
	resp.Body.Close()

	if result.Response == "False" {
		return nil, fmt.Errorf(result.Error)
	}

	return &result, nil
}

func downloadFile(URL, filename string) error {
	resp, err := http.Get(URL)
	if err != nil {
		resp.Body.Close()
		return err
	}
	if resp.StatusCode != http.StatusOK {
		resp.Body.Close()
		return fmt.Errorf("download failed: %s", resp.Status)
	}

	file, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		return err
	}

	return nil
}
