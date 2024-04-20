import { createStructure } from "/framework.js";
import { submitUsername } from "/ws/ws.js";


export default function registerView() {

    const root = document.getElementById('root'); // Ensure you have a div with id="root" in your HTML
    const structure = {
        tag: 'div',
        attr: ['id', 'entryForm', 'class', 'flex flex-col items-center justify-center min-h-screen bg-default'],
        children: [
            {tag: 'div',
                attr: ['class', 'flex flex-col items-center w-full max-w-lg bg-amber-300 shadow-md rounded px-16 pt-12 pb-8 mb-4'],
                children: [
            {
                tag: 'h1',
                attr: ['class', 'text-white text-3xl font-bold mb-4'],
                children: 'Enter your username'
            },
            {
                tag: 'input',
                attr: ['type', 'text', 'id', 'username', 'placeholder', 'Username', 'class', 'textbox-style'],
            },
            {
                tag: 'button',
                attr: ['onclick', () => submitUsername(), 'class', 'bg-green-700 hover:bg-green-600 text-white font-bold py-2 px-4 rounded'],
                children: 'Join Lobby'
            }
        ]}]
    };

    const content = createStructure(structure);
    root.innerHTML = '';  // Clear existing content
    root.appendChild(content);
}
