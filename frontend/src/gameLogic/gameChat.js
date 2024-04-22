import { createStructure, createChild } from "../framework.js";

export function createChatbox() {
    return createStructure({
        tag: 'div',
        attr: ['class', 'chatbox'],
        children: [
            {
                tag: 'h3',
                children: 'Messages'
            },
            {
                tag: 'div',
                attr: ['id', 'chatMessages']
            },
            {
                tag: 'div',
                attr: ['class', 'message-input-container'],
                children: [
                    {
                        tag: 'input',
                        attr: ['type', 'text', 'id', 'messageInput', 'placeholder', 'Type yuh message, mon!']
                    },
                    {
                        tag: 'button',
                        attr: ['id', 'sendButton'],
                        children: 'Send'
                    }
                ]
            }
        ]
    });
}

export function setupChat(ws) {
    // Get the DOM elements
    const messageInput = document.getElementById('messageInput');
    const sendButton = document.getElementById('sendButton');

    // Add event listener for send button
    sendButton.addEventListener('click', function () {
        const message = messageInput.value;
        if (message) {
            ws.send(JSON.stringify({ type: 'chatMessage', payload: message }));
            messageInput.value = '';
        }
    });

    // Add event listener for 'Enter' key in message input field
    messageInput.addEventListener('keydown', function (event) {
        if (event.key === 'Enter' && !event.shiftKey) {
            const message = messageInput.value;
            if (message) {
                ws.send(JSON.stringify({ type: 'chatMessage', payload: message }));
                messageInput.value = '';
            }
        }
    });
}

// Function to handle incoming chat messages
export function handleChatMessage(payload) {
    const chatMessages = document.getElementById('chatMessages');

    const structure = {
        tag: 'div',
        attr: ['class', 'message'],
        children: [
            {
                tag: 'p',
                attr: ['class', 'sender'],
                children: payload.Username
            },
            {
                tag: 'p',
                attr: ['class', 'message-content'],
                children: payload.Message
            },
            {
                tag: 'p',
                attr: ['class', 'time-sent'],
                children: payload.TimeSent
            }
        ]
    };

    // check if the message is from the current user by checking username in message and username in local storage
    // then change the class message to message-self
    if (payload.Username === sessionStorage.getItem('username')) {
        structure.attr[1] = 'message message-self';
    }

    chatMessages.appendChild(createStructure(structure));

    chatMessages.scrollTop = chatMessages.scrollHeight;
}

export function broadcastPlayerDisconnect(name) {
    const chatMessages = document.getElementById('chatMessages');

    const structure = {
        tag: 'div',
        attr: ['class', 'message'],
        children: [
            {
                tag: 'p',
                attr: ['class', 'player-left'],
                children: `${name} a lef' di game, ya hear?.`
            }
        ]
    };

    chatMessages.appendChild(createStructure(structure));

    chatMessages.scrollTop = chatMessages.scrollHeight;
}