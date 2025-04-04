# A GO server that uses the Spotify API to get a user's currently playing track
### Prerequisites:
  - A spotify account
  - Create an app in the Spotify account to obtain data crucial to running the app.

## Demo




### Here is a breakdown of the code :

## 1. Get Authourized
- Authourization begins in Auth2Flow.go file in the LoginHandler func
  ```go
  func LoginHandler(){
    //code
  }

## 2. Make Request
- After authourization, making api requests is now possible, this is done in the Getdata.go file inside the GetNowPlaying function.
```go
func GetNowPlaying(){
//code
}
```
- A struct called SongInfo is created to filter the json to just the parameters desired. Alternatively you can remove this step and sift through all the JSON data recieved.
```go
type SongInfo struct {
	SongName   string `json:"song_name"`
	ArtistName string `json:"artist_name"`
	AlbumArt   string `json:"album_art"`
	IsPlaying  bool   `json:"is_playing"`
}

```

## 3. Run Server
  
- Finally the server can be launched with previously mentioned functions as endpoints in the Auth2Flow.go file inside the Auth function.
```go
func Auth(){

 //code

}
```

- Run the program by running ` go run main.go Auth2Flow.go Getdata.go ` in the terminal

- Remember to run sequentially starting with the ` :8888/login ` endpoint then the ` :8888/getdata ` endpoint in the browser.
