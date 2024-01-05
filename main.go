package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	// GetNowPlaying("https://api.spotify.com/v1/me/player/currently-playing")

	// RequestToken()

	Authorize()

}

func GetNowPlaying(urls string) {
	token := "BQBTu_WH0dKt_1DHZNoD11fbF8O_AMd3e7p_dNlykyQFmborM6Hwal8Rs0PVEM6DzI7t294tzporNutCxvSRhmkWOxh4OJskN7yJZMNSrQd0V5MZN_4"

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

	fmt.Println(string(body))
}
