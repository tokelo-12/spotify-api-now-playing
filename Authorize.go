package main

import (
	"math/rand"
	"net/http"
	"net/url"
	"time"
)

const (
	clientID    = "0c7abc41f05d4c02ac74577d01798471"
	redirectURI = "http://localhost:3000."
	spotifyAuth = "https://accounts.spotify.com/authorize"
)

func Authorize() {

	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":3000", nil)
}

func generateRandomString(n int) string {
	rand.NewSource(time.Now().UnixNano())

	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {

	//set the state to generate random string for CSRF protection
	state := generateRandomString(16)

	//set scope
	scope := "user-read-currently-playing"

	params := url.Values{
		"response_type": {"code"},
		"client_id":     {clientID},
		"scope":         {scope},
		"redirect_uri":  {redirectURI},
		"state":         {state},
	}

	//attach params to spotify accounts auth
	authURL := spotifyAuth + "?" + params.Encode()

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

// import "github.com/gofiber/fiber/v2"
// app := fiber.New()

// app.Get("/", func(c *fiber.Ctx) error {
// 	return c.SendString("Hello, World 👋!")
// })

// app.Listen(":3000")
