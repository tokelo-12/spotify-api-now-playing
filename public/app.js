fetch("http://localhost:8888/callback")
    .then(response => response.json())
    .then(data => {
        // Process the received data
        console.log(data);
        // Use it in your HTML elements
        document.getElementsByClassName('track-name').innerText += data.song_name;
        document.getElementsByClassName('artist').innerText += data.artist_name;  
    })
    .catch(error => {
        console.error(`Error fetching data: ${error}`);
    });
