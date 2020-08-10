const assert = require('assert');

const isPalindrome = (s, startIndex, endIndex) => {
  let isPalindrome = false;
  let len = Math.floor(endIndex - startIndex);
  for (let i = 0; i < len; i++) {
    if (s[startIndex + i] !== s[endIndex - i]) {
      return false;
    }
  }

  return true;
};

/**
 * @param {string} s
 * @return {boolean}
 */
const validPalindrome = function (s) {
  let firstIndex = 0;
  let lastIndex = s.length - 1;

  while (firstIndex < lastIndex) {
    if (s[firstIndex] !== s[lastIndex]) {
      return (
        isPalindrome(s, firstIndex + 1, lastIndex) ||
        isPalindrome(s, firstIndex, lastIndex - 1)
      );
    }

    firstIndex++;
    lastIndex--;
  }

  return true;
};

assert.equal(validPalindrome('abc'), false);
assert.equal(validPalindrome('aba'), true);
assert.equal(validPalindrome('abca'), true);
assert.equal(
  validPalindrome(
    'aguokepatgbnvfqmgmlcupuufxoohdfpgjdmysgvhmvffcnqxjjxqncffvmhvgsymdjgpfdhooxfuupuculmgmqfvnbgtapekouga'
  ),
  true
);
