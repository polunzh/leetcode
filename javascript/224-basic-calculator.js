const assert = require('assert');
/**
 * @param {string} s
 * @return {number}
 */
var calculate = function(s) {
  const stack = [];
  let res = 0;
  let sign = 1;

  for (let i = 0; i < s.length; i++) {
    const n = s[i];
    if (n >= '0' && n <= '9') {
      let tmp = n;
      while (i + 1 < s.length && s[i + 1] >= '0' && s[i + 1] <= '9') {
        tmp += s[++i];
      }

      res += sign * Number(tmp);
    } else if (n === '+') {
      sign = 1;
    } else if (n === '-') {
      sign = -1;
    } else if (n === '(') {
      stack.push(res);
      stack.push(sign);
      res = 0;
      sign = 1;
    } else if (n === ')') {
      const s = stack.pop();
      const r = stack.pop();
      res = r + s * res;
      sign = 1;
    }
  }

  return res;
};

assert.equal(3 + 1, calculate('3+1'));
assert.equal(3 + 1, calculate('(3+1)'));
assert.equal(3 + 1 - 1, calculate('(3+1) - 1'));
assert.equal(3 + 1 - 12 - (1 + 1), calculate('(3+1) - 12 - (1+1)'));
assert.equal(3 + 1 - 12 - (1 + 1) + 1, calculate('(3+1) - 12 - (1+1) + 1'));
