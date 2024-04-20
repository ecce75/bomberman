var state = {
    todo: [],
    selectedBtn: 1,
    page: '/#',
    filter: '',
    hideInfo: false,
    hideClear: false,
    allChecked: false,
};

function saveState(currentState) {
    localStorage.setItem('state', JSON.stringify(currentState));
}

function loadState() {
    let data = localStorage.getItem('state');
    if (data) {
        state = JSON.parse(data);
    }
}

function createChild(parent, child) {
    if (child instanceof Node) {
        parent.appendChild(child);
    } else if (typeof child === 'string' || child instanceof String) {
        parent.appendChild(document.createTextNode(child));
    } else if (typeof child === 'object') {
        // Assume it's a structure object and needs to be created
        const childElement = createStructure(child);
        parent.appendChild(childElement);
    } else {
        console.log('Unsupported child type:', child);
    }
}

function removeChild(parent, child) {
    parent.removeChild(child);
}
function createStructure(structure) {
    let parent = document.createElement(structure.tag);
    if ('attr' in structure) {
        setAttributes(parent, structure.attr);
    }
    if ('style' in structure) {
        setStyles(parent, structure.style); // Call setStyles function if style attribute is present
    }
    if ('children' in structure) {
        if (Array.isArray(structure.children)) {
            for (const child of structure.children) {
                createChild(parent, child);
            }
        } else {
            createChild(parent, structure.children);
        }
    }

    return parent;
}

function setAttributes(element, attributes) {
    for (let i = 0; i < attributes.length; i += 2) {
        let key = attributes[i];
        let value = attributes[i + 1];
        if (value === null || value === undefined) {
            continue;
        }
        if (key.startsWith('on')) {
            const eventName = key.slice(2).toLowerCase();
            console.log(eventName, value, key)
            element.addEventListener(eventName, value);
        } else {
            element.setAttribute(key, value);
        }
    }
}

function setStyles(element, styles) {
    for (let i = 0; i < styles.length; i += 2) {
        let key = styles[i];
        let value = styles[i + 1];
        console.log(key, value)
        if (value === null || value === undefined) {
            continue;
        }
        element.style[key] = value;
    }
}

function getParent(element, num = 2) {
    for (let i = 0; i < num; i++) {
        element = element.parentElement;
    }
    return element;
}

function redirect(url) {
    window.history.pushState(null, null, url);
}

function addEvent(eventType, element, callback) {
    element.addEventListener(eventType, callback);
}

function removeEvent(eventType, element) {
    element.removeEventListener(eventType);
}

export { state, saveState, loadState, createChild, removeChild, createStructure, setAttributes, setStyles, getParent, redirect, addEvent, removeEvent}