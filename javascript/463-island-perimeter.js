/**
 * @param {number[][]} grid
 * @return {number}
 */
var islandPerimeter = function(grid) {
  let perimeter = 0;
  for (let i = 0; i < grid.length; i++) {
    for (let j = 0; j < grid[i].length; j++) {
      if (grid[i][j] === 1) {
        // up
        if (j === 0 || grid[i][j - 1] !== 1) {
          perimeter++;
        }

        // right
        if (i === grid.length - 1 || grid[i + 1][j] !== 1) {
          perimeter++;
        }

        // down
        if (j === grid[i].length - 1 || grid[i][j + 1] !== 1) {
          perimeter++;
        }

        // left
        if (i === 0 || grid[i - 1][j] !== 1) {
          perimeter++;
        }
      }
    }
  }

  return perimeter;
};

console.log(
  islandPerimeter([[0, 0, 0, 0], [1, 1, 1, 0], [0, 1, 0, 0], [1, 1, 0, 0]])
);
