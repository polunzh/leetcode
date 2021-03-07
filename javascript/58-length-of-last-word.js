/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function (s) {
  let current = 0;
  for (let i = s.length - 1; i >= 0; i--) {
    if (s[i] === ' ') {
      if (current === 0) {
        continue;
      } else {
        break;
      }
    } else {
      if (s[i] !== ' ') {
        current++;
      }
    }
  }

  return current;
};

console.log(lengthOfLastWord(' i   '));
console.log(lengthOfLastWord('Hello   World'));
