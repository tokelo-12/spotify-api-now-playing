fetch("http://localhost:8888/getdata")
    .then(response => response.json())
    .then(data => {
        // Process the received data
        console.log(data);
        // Use it in your HTML elements
        document.querySelector(".track-name").textContent = data.song_name;
        document.querySelector('.artist').textContent = data.artist_name ;  

        document.querySelector('.album-art').innerHTML = '<img src="' + data.album_art + '" alt="Album Art">';
    })
    .catch(error => {
        console.error(`Error fetching data: ${error}`);
    });
