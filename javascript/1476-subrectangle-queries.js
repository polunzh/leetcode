/**
 * @param {number[][]} rectangle
 */
const SubrectangleQueries = function (rectangle) {
  this.rectangle = rectangle;
};

/**
 * @param {number} row1
 * @param {number} col1
 * @param {number} row2
 * @param {number} col2
 * @param {number} newValue
 * @return {void}
 */
SubrectangleQueries.prototype.updateSubrectangle = function (
  row1,
  col1,
  row2,
  col2,
  newValue
) {
  for (let row = row1; row <= row2; row++) {
    for (let col = col1; col <= col2; col++) {
      this.rectangle[row][col] = newValue;
    }
  }
};

/**
 * @param {number} row
 * @param {number} col
 * @return {number}
 */
SubrectangleQueries.prototype.getValue = function (row, col) {
  return this.rectangle[row][col];
};

/**
 * Your SubrectangleQueries object will be instantiated and called as such:
 * var obj = new SubrectangleQueries(rectangle)
 * obj.updateSubrectangle(row1,col1,row2,col2,newValue)
 * var param_2 = obj.getValue(row,col)
 */

const subrectangleQueries = new SubrectangleQueries([[1,1,1],[2,2,2],[3,3,3]]);
console.log(subrectangleQueries.getValue(0, 0)); // return 1
subrectangleQueries.updateSubrectangle(0, 0, 2, 2, 100);
console.log(subrectangleQueries.getValue(0, 0)); // return 100
console.log(subrectangleQueries.getValue(2, 2)); // return 100
subrectangleQueries.updateSubrectangle(1, 1, 2, 2, 20);
console.log(subrectangleQueries.getValue(2, 2)); // return 20