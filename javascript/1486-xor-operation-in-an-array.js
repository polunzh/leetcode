/**
 * @param {number} n
 * @param {number} start
 * @return {number}
 */
const xorOperation = function (n, start) {
  let result = 0;

  for (let i = 0; i < n; i++) {
    result ^= start + 2 * i;
  }

  return result;
};

console.log(xorOperation(5, 0));
console.log(xorOperation(1, 7));
console.log(xorOperation(10, 5));
