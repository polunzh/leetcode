/**
 * @param {number} numRows
 * @return {number[][]}
 */
const generate = function (numRows) {
  if (numRows < 1) {
    return [];
  }

  const result = [[1]];
  for (let i = 1; i < numRows; i++) {
    const tmpArray = [1];
    for (let j = 0; j < result[result.length - 1].length - 1; j++) {
      tmpArray.push(
        result[result.length - 1][j] + result[result.length - 1][j + 1]
      );
    }

    tmpArray.push(1);
    result.push(tmpArray);
  }

  return result;
};

console.log(generate(1));
console.log(generate(5));
