package main

import (
	"crypto/rand"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"strings"

	"github.com/rs/cors"
)

const (
	clientID     = ""
	clientSecret = ""
	redirectURI  = "http://localhost:8888/callback"
	stateKey     = "spotify_auth_state"
	spotifyAuth  = "https://accounts.spotify.com/authorize"
)

//This function generates a random string of a specified length using cryptographic random bytes.

func generateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

// This function is the handler for the /login route. It generates a random state, sets it as a cookie, defines the Spotify API scope, and redirects the user to the Spotify authorization page.

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateRandomString(16)
	http.SetCookie(w, &http.Cookie{
		Name:  stateKey,
		Value: state,
		Path:  "/",
	})
	scope := "user-read-currently-playing"

	params := url.Values{
		"response_type": {"code"},
		"client_id":     {clientID},
		"scope":         {scope},
		"redirect_uri":  {redirectURI},
		"state":         {state},
	}

	authURL := spotifyAuth + "?" + params.Encode()

	http.Redirect(w, r, authURL, http.StatusTemporaryRedirect)
}

var myToken string

func callBackHandler(w http.ResponseWriter, r *http.Request) {

	//Extract code, state, and storedState from request

	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")
	storedState, err := r.Cookie(stateKey)
	if err != nil || state == "" || state != storedState.Value {
		http.Redirect(w, r, "/#"+url.Values{"error": {"state_mismatch"}}.Encode(), http.StatusTemporaryRedirect)
		return
	}

	http.SetCookie(w, &http.Cookie{
		Name:   stateKey,
		Value:  "",
		Path:   "/",
		MaxAge: -1,
	})

	authOptions := url.Values{
		"code":         {code},
		"redirect_uri": {redirectURI},
		"grant_type":   {"authorization_code"},
	}

	// Create an HTTP request to exchange the code for tokens

	req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", strings.NewReader(authOptions.Encode()))

	if err != nil {
		http.Redirect(w, r, "/#"+url.Values{"error": {"internal_error"}}.Encode(), http.StatusTemporaryRedirect)
		return
	}

	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	req.Header.Set("Authorization", "Basic "+base64.StdEncoding.EncodeToString([]byte(clientID+":"+clientSecret)))

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		http.Redirect(w, r, "/#"+url.Values{"error": {"internal_error"}}.Encode(), http.StatusTemporaryRedirect)
		return
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		http.Redirect(w, r, "/#"+url.Values{"error": {"invalid_token"}}.Encode(), http.StatusTemporaryRedirect)
		return
	}

	// Parse the response body to get the access and refresh tokens

	var tokenResponse map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&tokenResponse)
	if err != nil {
		http.Redirect(w, r, "/#"+url.Values{"error": {"invalid_token"}}.Encode(), http.StatusTemporaryRedirect)
		return
	}

	myToken = tokenResponse["access_token"].(string)
	// refreshToken := tokenResponse["refresh_token"].(string)

	// Handle the access token and make requests as needed

	// GetNowPlaying(w, r, accessToken)

	//pass the tokens to the browser
	// http.Redirect(w, r, "/#"+url.Values{"access_token": {accessToken}, "refresh_token": {refreshToken}}.Encode(), http.StatusTemporaryRedirect)

}

// type corsHandler struct {
// 	handler http.Handler
// }

// func (c *corsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
// 	c.handler.ServeHTTP(w, r)
// }

func Auth() {

	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // Replace with your frontend origin
		AllowedMethods:   []string{"HEAD", "GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type"},
		AllowCredentials: true,
		Debug:            false, // Set to false in production
	})

	http.HandleFunc("/login", LoginHandler)

	http.HandleFunc("/callback", callBackHandler)
	http.Handle("/getdata", c.Handler(http.HandlerFunc(GetNowPlaying)))

	http.ListenAndServe(":8888", nil)

	fmt.Println("Listening on :8888")
}
