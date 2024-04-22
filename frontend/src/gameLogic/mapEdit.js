export function normalizeField(coordinates) {
    var field = document.querySelector(`.cell[style*="grid-area: ${coordinates.Y + 1} / ${coordinates.X + 1}"]`);
    field.className = 'cell';
}