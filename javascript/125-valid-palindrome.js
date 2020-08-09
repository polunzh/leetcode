const assert = require('assert');

const isValidChar = (c) =>
  (c.charCodeAt() >= 48 && c.charCodeAt() <= 57) ||
  (c.charCodeAt() >= 65 && c.charCodeAt() <= 90) ||
  (c.charCodeAt() >= 97 && c.charCodeAt() <= 122);

/**
 * @param {string} s
 * @return {boolean}
 */
const isPalindrome = function (s) {
  let firstIndex = 0;
  let secondIndex = s.length - 1;

  while (firstIndex <= secondIndex) {
    if (!isValidChar(s[firstIndex])) {
      firstIndex++;
      continue;
    }

    while (!isValidChar(s[secondIndex])) {
      if (secondIndex < firstIndex) {
        return false;
      }

      secondIndex--;
    }

    if (s[firstIndex].toLowerCase() !== s[secondIndex].toLowerCase()) {
      return false;
    }

    firstIndex++;
    secondIndex--;
  }

  return true;
};

assert.equal(isPalindrome('A man, a plan, a canal: Panama'), true);
assert.equal(isPalindrome('race a car'), false);
