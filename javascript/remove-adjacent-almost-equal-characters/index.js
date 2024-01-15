/**
 * @param {string} word
 * @return {number}
 */
const removeAlmostEqualCharacters = function (word) {
  let result = 0;
  if (word.length < 1) {
    return result;
  }

  let header = 0;
  for (let i = 1; i < word.length; i++) {
    if (i < word.length - 1) {
      if (
        Math.abs(word[i].charCodeAt() - word[i - 1].charCodeAt()) <= 1 &&
        Math.abs(word[i].charCodeAt() - word[i + 1].charCodeAt()) <= 1
      ) {
        result++;
        header = i + 1;
        i += 1;
      } else if (
        Math.abs(word[i].charCodeAt() - word[i - 1].charCodeAt()) <= 1 &&
        header === i - 1
      ) {
        header = i + 1;
        i++;
        result++;
      } else {
        header += 1;
      }
    } else {
      if (Math.abs(word[i].charCodeAt() - word[i - 1].charCodeAt()) <= 1) {
        result++;
      }
    }
  }

  return result;
};

module.exports = removeAlmostEqualCharacters;
