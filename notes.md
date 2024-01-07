
### Spotify Client Credential Auth flow


```go

package main

import (
    "bytes"
    "fmt"
    "io/ioutil"
    "net/http"
    "net/url"
)

func main() {
    clientID := "your-client-id"
    clientSecret := "your-client-secret"

    data := url.Values{}
    data.Set("grant_type", "client_credentials")
    data.Set("client_id", clientID)
    data.Set("client_secret", clientSecret)

    req, err := http.NewRequest("POST", "https://accounts.spotify.com/api/token", bytes.NewBufferString(data.Encode()))
    if err != nil {
        panic(err)
    }

    req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        panic(err)
    }

    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        panic(err)
    }

    fmt.Println(string(body))
}

```

### Authorization Code AUTH flow

```golang

package main

import (
    "fmt"
    "net/http"
    "net/url"
)

const (
    clientID     = "YOUR_CLIENT_ID"
    redirectURI  = "http://localhost:8888/callback"
    spotifyAuth  = "https://accounts.spotify.com/authorize"
)

func main() {
    http.HandleFunc("/login", loginHandler)
    http.ListenAndServe(":8888", nil)
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
    state := generateRandomString(16) // Implement this function for CSRF protection
    scope := "user-read-private user-read-email"

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

```

### Authorization Code oAuth2 flow to authenticate against the Spotify Accounts

```golang
package main

import (
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

var (
	clientID     = "yourClientIDGoesHere"
	clientSecret = "YourSecretIDGoesHere"
	redirectURI  = "http://localhost:8888/callback"
	stateKey     = "spotify_auth_state"
)

func generateRandomString(length int) string {
	b := make([]byte, length)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)[:length]
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	state := generateRandomString(16)
	http.SetCookie(w, &http.Cookie{
		Name:  stateKey,
		Value: state,
		Path:  "/",
	})
	scope := "user-read-private user-read-email"
	redirectURL := fmt.Sprintf("https://accounts.spotify.com/authorize?response_type=code&client_id=%s&scope=%s&redirect_uri=%s&state=%s",
		clientID, url.QueryEscape(scope), url.QueryEscape(redirectURI), url.QueryEscape(state))
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
}

func callbackHandler(w http.ResponseWriter, r *http.Request) {
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
		"code":          {code},
		"redirect_uri":  {redirectURI},
		"grant_type":    {"authorization_code"},
	}

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

	// Handle the access token and make requests as needed
	// (similar to the Node.js code)

	http.Redirect(w, r, "/#"+url.Values{"access_token": {accessToken}, "refresh_token": {refreshToken}}.Encode(), http.StatusTemporaryRedirect)
}

func main() {
	http.HandleFunc("/login", loginHandler)
	http.HandleFunc("/callback", callbackHandler)
	fmt.Println("Listening on :8888")
	http.ListenAndServe(":8888", nil)
}
```





