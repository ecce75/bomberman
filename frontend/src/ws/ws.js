import {lobbyView} from "/views/lobbyView.js";



function setupWebSocket() {
    const ws = new WebSocket('ws://localhost:8080/ws'); // Adjust this URL to your server
    ws.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        switch(msg.type) {
            case 'updateCounter':
                document.getElementById('playerCount').textContent = msg.payload.toString();
                break;
            case 'updateTime':
                document.getElementById('playerCount').textContent = msg.payload.toString();
                break;
            case "invalidUsername":
                alert("Username already taken")
                window.reload()
            // Handle other messages
        }
    };

    return ws;
}

export function submitUsername() {
    const username = document.getElementById('username').value.trim();
    if (!username) {
        alert('Please enter a username.');
        return;
    }
    const ws = setupWebSocket();
    ws.onopen = function() {
        ws.send(JSON.stringify({ type: 'setUsername', payload: username }));
        sessionStorage.setItem("username", username)
        lobbyView();
    };
}


