
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





