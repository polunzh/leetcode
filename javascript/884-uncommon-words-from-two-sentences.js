/**
 * @param {string} A
 * @param {string} B
 * @return {string[]}
 */
var uncommonFromSentences = function(A, B) {
  const tmp = {};

  const arr = A.split(/\s/).concat(B.split(/\s/));
  const result = [];
  arr.forEach(item => {
    if (tmp[item]) {
      tmp[item]++;
    } else {
      tmp[item] = 1;
    }
  });

  Object.keys(tmp).forEach(x => {
    if (tmp[x] === 1) {
      result.push(x);
    }
  });

  return result;
};

console.log(uncommonFromSentences('this apple is sweet', 'this apple is sour'));
console.log(uncommonFromSentences('apple apple', 'banana'));
