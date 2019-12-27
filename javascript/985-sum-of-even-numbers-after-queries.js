/**
 * @param {number[]} A
 * @param {number[][]} queries
 * @return {number[]}
 */
var sumEvenAfterQueries = function(A, queries) {
  const result = [];
  for (var i = 0; i < A.length; i++) {
    var val = queries[i][0];
    var index = queries[i][1];
    var tmp = A[index];
    A[index] += val;

    if (result.length === 0) {
      result.push(A.filter(x => x % 2 === 0).reduce((acc, x) => (acc += x), 0));
    } else if (A[index] % 2 === 0 || A[index] % 2 === -0) {
      if (tmp % 2 !== 0) {
        result.push(result[result.length - 1] + A[index]);
      } else {
        result.push(result[result.length - 1] + val);
      }
    } else {
      if (tmp % 2 === 0) {
        result.push(result[result.length - 1] - tmp);
      } else {
        result.push(result[result.length - 1]);
      }
    }
  }

  return result;
};

console.log(
  sumEvenAfterQueries(
    [1, 2, 3, 4],
    [
      [1, 0],
      [-3, 1],
      [-4, 0],
      [2, 3],
    ]
  )
);
console.log(
  sumEvenAfterQueries(
    [5, 1],
    [
      [0, 1],
      [4, 0],
    ]
  )
);
console.log(
  sumEvenAfterQueries(
    [5, 5, 4],
    [
      [0, 1],
      [1, 0],
      [4, 1],
    ]
  )
);
