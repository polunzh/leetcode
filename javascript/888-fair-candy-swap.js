const assert = require('assert');
/**
 * @param {number[]} A
 * @param {number[]} B
 * @return {number[]}
 */
var fairCandySwap = function(A, B) {
  let sumA = 0;
  let sumB = 0;
  for (let i = 0; i < A.length; i++) {
    sumA += A[i];
  }

  for (let i = 0; i < B.length; i++) {
    sumB += B[i];
  }

  const bSet = new Set(B);
  for (let i = 0; i < A.length; i++) {
    const t = (sumB - sumA + 2 * A[i]) / 2;
    if (bSet.has(t)) {
      return [A[i], t];
    }
  }

  return [];
};

assert.deepEqual(fairCandySwap([1, 1], [2, 2]), [1, 2]);
assert.deepEqual(fairCandySwap([1, 2], [2, 3]), [1, 2]);
assert.deepEqual(fairCandySwap([2], [1, 3]), [2, 3]);
assert.deepEqual(fairCandySwap([1, 2, 5], [2, 4]), [5, 4]);
