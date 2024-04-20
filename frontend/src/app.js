// Example in app.js
import './views/registerView.js';
import './views/lobbyView.js';
import './views/gameView.js';
import {createChild} from "./framework.js";
import './ws/ws.js';
import registerView from "./views/registerView.js";

// app.js
document.addEventListener('DOMContentLoaded', function() {
    initializeApp();
});


function initializeApp() {
    const root = document.getElementById('root');
    root.innerHTML = ''; // Clear previous content if any
    createChild(root, registerView());
}
