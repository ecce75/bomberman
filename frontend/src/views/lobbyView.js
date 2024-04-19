function lobbyView() {
    const structure = {
        tag: 'div',
        attr: ['id', 'lobby'],
        children: [
            { tag: 'h1', children: 'Lobby' },
            { tag: 'p', children: 'Waiting for more players...' },
            {
                tag: 'p',
                children: [
                    'Players in lobby: ',
                    { tag: 'span', attr: ['id', 'playerCount'], children: '0' }
                ]
            }
        ]
    };

    return createStructure(structure);
}
