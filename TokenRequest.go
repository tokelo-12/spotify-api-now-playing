package main

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func RequestToken() {
	clientID := ""
	clientSecret := ""

	//prepare form data
	var data = url.Values{}

	data.Set("client_id", clientID)
	data.Set("client_secret", clientSecret)
	data.Set("grant_type", "client_credentials")

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))

	if err != nil {
		panic(err)
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	//send the request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	//runs resp.Body.Close after the parent function returns, it closes the response body to free up resources
	defer resp.Body.Close()

	//read http response body with io.ReadAll
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(body))

}
