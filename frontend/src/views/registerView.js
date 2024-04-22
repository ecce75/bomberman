import { createStructure } from "/framework.js";
import { submitUsername } from "/ws/ws.js";
import {addEvent} from "../framework.js";


export default function registerView() {

    addEvent('keydown', document, function(event) {
        if (event.key === 'Enter') {
            submitUsername();
        }
    })

    const root = document.getElementById('root'); // Ensure you have a div with id="root" in your HTML
    const structure = {
        tag: 'div',
        attr: ['id', 'entryForm', 'class', 'flex flex-col items-center justify-center min-h-screen bg-default'],
        children: [
            {tag: 'div',
                attr: ['class', 'flex flex-col items-center w-full max-w-lg bg-gradient-to-t from-amber-500 to-red-600 shadow-md rounded-lg px-16 pt-12 pb-8 mb-4'],
                children: [
            {
                tag: 'h1',
                attr: ['class', 'text-white text-3xl font-bold mb-4'],
                children: 'Who yuh be, mon?'
            },
            {
                tag: 'input',
                attr: ['type', 'text', 'id', 'username', 'placeholder', 'Username', 'class', 'textbox-style'],
            },
            {
                tag: 'button',
                attr: ['onclick', () => submitUsername(), 'class', 'bg-green-700 hover:bg-green-600 text-white font-bold py-2 px-4 rounded'],
                children: 'Join in, mon!'
            }
        ]}]
    };

    const content = createStructure(structure);
    root.innerHTML = '';  // Clear existing content
    root.appendChild(content);
}
