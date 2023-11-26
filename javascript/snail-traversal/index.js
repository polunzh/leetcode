/**
 * @param {number} rowsCount
 * @param {number} colsCount
 * @return {Array<Array<number>>}
 */
Array.prototype.snail = function (rowsCount, colsCount) {
  if (this.length !== rowsCount * colsCount) {
    return [];
  }

  const arr = new Array(rowsCount).fill([]);
  arr.forEach((_, i) => (arr[i] = []));

  for (let i = 0; i < this.length; i += rowsCount) {
    const index = i / rowsCount;
    if (index % 2 === 0) {
      for (let j = 0; j < arr.length; j++) {
        arr[j][index] = this[i + j];
      }
    } else {
      for (let j = arr.length - 1; j >= 0; j--) {
        arr[j][index] = this[i + (arr.length - j - 1)];
      }
    }
  }

  return arr;
};
