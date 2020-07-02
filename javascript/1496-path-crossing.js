/**
 * @param {string} path
 * @return {boolean}
 */
const isPathCrossing = function (path) {
  let x = 0;
  let y = 0;
  const xMap = new Map();
  const yMap = new Map();
  xMap.set('E', 1);
  xMap.set('W', -1);

  yMap.set('N', 1);
  yMap.set('S', -1);

  const visitedSet = new Set();
  visitedSet.add(`${x},${y}`);

  for (let i = 0; i < path.length; i++) {
    if (xMap.has(path[i])) {
      x += xMap.get(path[i]);
    } else {
      y += yMap.get(path[i]);
    }

    const point = `${x},${y}`;
    if (visitedSet.has(point)) {
      return true;
    }

    visitedSet.add(point);
  }

  return false;
};

console.log(isPathCrossing('NES'));
console.log(isPathCrossing('NESWW'));
