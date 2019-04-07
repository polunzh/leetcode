/**
 * @param {string} A
 * @param {string} B
 * @return {boolean}
 */
var buddyStrings = function(A, B) {
  if (A.length != B.length) {
    return false;
  }

  if (!A || !B) {
    return false;
  }

  if (A.length == 1) {
    return false;
  }

  if (A === B) {
    const s = new Set(A);
    return s.size === A.length ? false : true;
  }

  let tmp = [];
  for (let i = 0; i < A.length; i++) {
    if (A[i] !== B[i]) {
      tmp.push(...[A[i], B[i]]);
    }

    if (tmp.length > 4) {
      return false;
    }
  }

  if (tmp.length === 0) {
    return true;
  }

  if (tmp.length !== 4) {
    return false;
  }

  return tmp[0] === tmp[3] && tmp[1] === tmp[2];
};

console.log(buddyStrings('ab', 'ab'));
console.log(buddyStrings('aab', 'aab'));
console.log(buddyStrings('aba', 'aab'));
