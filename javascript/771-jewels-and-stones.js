/**
 * @param {string} J
 * @param {string} S
 * @return {number}
 */
var numJewelsInStones = function(J, S) {
  const map = J.split('').reduce((acc, cur) => {
    acc[cur] = 1;
    return acc;
  }, {});
  let result = 0;
  for (let i = 0; i < S.length; i++) {
    if (map[S[i]]) {
      result++;
    }
  }

  return result;
};

const assert = require('assert');
assert.equal(3, numJewelsInStones('aA', 'aAAbbbb'));
assert.equal(0, numJewelsInStones('z', 'ZZ'));
