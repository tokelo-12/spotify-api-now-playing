package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// GetNowPlaying("https://api.spotify.com/v1/me/player/currently-playing")

	// RequestToken()

	Auth()

}

func GetNowPlaying(urls string, token string) {

	req, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		panic(err)
	}

	//set request header
	req.Header.Set("Authorization", "Bearer "+token)

	//send request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//read response
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Print(string(body))
}
