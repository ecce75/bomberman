function setupWebSocket() {
    const ws = new WebSocket('ws://localhost:8080/ws'); // Adjust this URL to your server
    ws.onmessage = function(event) {
        const msg = JSON.parse(event.data);
        switch(msg.type) {
            case 'updateCounter':
                document.getElementById('playerCount').textContent = msg.payload.toString();
                break;
            // Handle other messages
        }
    };

    return ws;
}

function submitUsername() {
    const username = document.getElementById('username').value.trim();
    if (!username) {
        alert('Please enter a username.');
        return;
    }
    state.username = username;
    const ws = setupWebSocket();
    ws.onopen = function() {
        ws.send(JSON.stringify({ type: 'setUsername', payload: username }));
        switchViewToLobby();
    };
}

function switchViewToLobby() {
    const root = document.getElementById('root'); // Assuming your HTML has a div with id="root"
    removeChild(root, document.getElementById('entryForm'));
    createChild(root, lobbyView());
}
