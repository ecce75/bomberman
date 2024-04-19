import {createStructure} from "/framework.js";

export default function registerView() {
    const root = document.getElementById('root'); // Ensure you have a div with id="app-root" in your HTML
    const structure = {
        tag: 'div',
        attr: ['id', 'entryForm'],
        children: [
            {
                tag: 'h1',
                children: 'Enter your username'
            },
            {
                tag: 'input',
                attr: ['type', 'text', 'id', 'username', 'placeholder', 'Username']
            },
            {
                tag: 'button',
                attr: ['onclick', 'submitUsername()'],
                children: 'Join Lobby'
            }
        ]
    };

    const content = createStructure(structure);
    root.appendChild(content);
}


