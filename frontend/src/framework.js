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

/**
 * Appends a child element to a parent element.
 * 
 * @param {Node} parent - The parent element to append the child to.
 * @param {Node|string|object} child - The child element to append. It can be a Node, a string representing text content, or an object representing a structure.
 */
function createChild(parent, child) {
    if (child instanceof Node) {
        parent.appendChild(child);
    } else if (typeof child === 'string' || child instanceof String) {
        parent.appendChild(document.createTextNode(child));
    } else if (typeof child === 'object') {
        // Assume it's a structure object and needs to be created
        const childElement = createStructure(child);
        parent.appendChild(childElement);
    }
}

function removeChild(parent, child) {
    parent.removeChild(child);
}


/**
 * Creates a DOM structure based on the provided configuration object.
 *
 * @param {Object} structure - The configuration object describing the structure.
 * @param {string} structure.tag - The HTML tag name for the parent element.
 * @param {Object} [structure.attr] - The attributes to be set on the parent element.
 * @param {Object} [structure.style] - The styles to be applied to the parent element.
 * @param {Array|Object} [structure.children] - The child elements to be appended to the parent element.
 * @returns {HTMLElement} - The created parent element.
 */
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

/**
 * Sets attributes on an HTML element.
 *
 * @param {HTMLElement} element - The HTML element to set attributes on.
 * @param {Array<string>} attributes - An array of attribute key-value pairs.
 */
function setAttributes(element, attributes) {
    for (let i = 0; i < attributes.length; i += 2) {
        let key = attributes[i];
        let value = attributes[i + 1];
        if (value === null || value === undefined) {
            continue;
        }
        if (key.startsWith('on')) {
            const eventName = key.slice(2).toLowerCase();
            element.addEventListener(eventName, value);
        } else {
            element.setAttribute(key, value);
        }
    }
}

/**
 * Sets the styles of an element based on the provided key-value pairs.
 *
 * @param {HTMLElement} element - The element to apply the styles to.
 * @param {Array} styles - An array of key-value pairs representing the styles to be set.
 */
function setStyles(element, styles) {
    for (let i = 0; i < styles.length; i += 2) {
        let key = styles[i];
        let value = styles[i + 1];
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