import {createStructure} from "/framework.js";

export function lobbyView() {
    const root = document.getElementById('root'); // Ensure you have a div with id="root" in your HTML

    const structure = {
        tag: 'div',
        attr: ['class', 'flex flex-col items-center justify-center min-h-screen bg-default'],
        children: [{
        tag: 'div',
        attr: ['id', 'lobby'],
        children: [
            { tag: 'h1', children: 'Lobby' },
            { tag: 'p', children: 'Waiting for more players...' },
            { tag: 'p', children: ['Game will start in ... ', { tag: 'span', attr: ['id', 'countdown'], children: '' }, ' seconds.' ]},
            {
                tag: 'p',
                children: [
                    'Players in lobby: ',
                    { tag: 'span', attr: ['id', 'playerCount'], children: '0' }
                ]
            }
        ]}]
    };

    const content = createStructure(structure);
    root.innerHTML = '';  // Clear existing content
    root.appendChild(content);
}
