package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type ExternalUrls struct {
	Spotify string `json:"spotify"`
}

type Artist struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	ID           string       `json:"id"`
	Name         string       `json:"name"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type Album struct {
	AlbumType        string       `json:"album_type"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	Images           []struct {
		Height int    `json:"height"`
		URL    string `json:"url"`
		Width  int    `json:"width"`
	} `json:"images"`
	Name                 string `json:"name"`
	ReleaseDate          string `json:"release_date"`
	ReleaseDatePrecision string `json:"release_date_precision"`
	TotalTracks          int    `json:"total_tracks"`
	Type                 string `json:"type"`
	URI                  string `json:"uri"`
}

type ExternalIds struct {
	Isrc string `json:"isrc"`
}

type Item struct {
	Album            Album        `json:"album"`
	Artists          []Artist     `json:"artists"`
	AvailableMarkets []string     `json:"available_markets"`
	DiscNumber       int          `json:"disc_number"`
	DurationMs       int          `json:"duration_ms"`
	Explicit         bool         `json:"explicit"`
	ExternalIds      ExternalIds  `json:"external_ids"`
	ExternalUrls     ExternalUrls `json:"external_urls"`
	Href             string       `json:"href"`
	ID               string       `json:"id"`
	IsLocal          bool         `json:"is_local"`
	Name             string       `json:"name"`
	Popularity       int          `json:"popularity"`
	PreviewUrl       string       `json:"preview_url"`
	TrackNumber      int          `json:"track_number"`
	Type             string       `json:"type"`
	URI              string       `json:"uri"`
}

type Actions struct {
	Disallows struct {
		Resuming bool `json:"resuming"`
	} `json:"disallows"`
}

type Context struct {
	ExternalUrls ExternalUrls `json:"external_urls"`
	Href         string       `json:"href"`
	Type         string       `json:"type"`
	URI          string       `json:"uri"`
}

type NowPlaying struct {
	Timestamp            int64   `json:"timestamp"`
	Context              Context `json:"context"`
	ProgressMs           int     `json:"progress_ms"`
	Item                 Item    `json:"item"`
	CurrentlyPlayingType string  `json:"currently_playing_type"`
	Actions              Actions `json:"actions"`
	IsPlaying            bool    `json:"is_playing"`
}

type SongInfo struct {
	SongName   string `json:"song_name"`
	ArtistName string `json:"artist_name"`
	AlbumArt   string `json:"album_art"`
}

func GetNowPlaying(w http.ResponseWriter, r *http.Request, token string) {

	urls := "https://api.spotify.com/v1/me/player/currently-playing"

	req, err := http.NewRequest("GET", urls, nil)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	//set request header
	req.Header.Set("Authorization", "Bearer "+token)

	//send request
	client := &http.Client{}

	resp, err := client.Do(req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	defer resp.Body.Close()

	//read response
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// fmt.Print(string(body))

	// Set response headers
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json")

	var data NowPlaying
	error := json.Unmarshal(body, &data)
	if error != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	fmt.Printf("Song Name: %s\n", data.Item.Name)
	fmt.Printf("Artist Name: %s\n", data.Item.Artists[0].Name)
	fmt.Printf("Album Art: %s\n", data.Item.Album.Images[0].URL)

	songInfo := SongInfo{
		SongName:   data.Item.Name,
		ArtistName: data.Item.Artists[0].Name,
		AlbumArt:   data.Item.Album.Images[0].URL,
	}

	// Send JSON response to the frontend if all data should be sent
	// json.NewEncoder(w).Encode(data)

	json.NewEncoder(w).Encode(songInfo)
}
