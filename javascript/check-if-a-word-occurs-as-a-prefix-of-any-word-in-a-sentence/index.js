/**
 * @param {string} sentence
 * @param {string} searchWord
 * @return {number}
 */
const isPrefixOfWord = function (sentence, searchWord) {
  let wordIndex = 1;

  for (let i = 0; i < sentence.length; ) {
    // compare
    let matched = true;
    for (let j = 0; j < searchWord.length; j++) {
      if (sentence[i + j] !== searchWord[j]) {
        matched = false;
        break;
      }
    }

    if (matched) {
      return wordIndex;
    }

    // next word
    let k = i;
    for (; k <= sentence.length && sentence[k] !== ' '; k++) {}

    if (k >= sentence.length) {
      return -1;
    }

    i = k + 1;
    wordIndex++;
  }

  return -1;
};

module.exports = isPrefixOfWord;
