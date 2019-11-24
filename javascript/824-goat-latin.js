/**
 * @param {string} S
 * @return {string}
 */
var toGoatLatin = function(S) {
  if (!S) {
    return S;
  }

  var arr = S.split(' ');
  return arr
    .map((word, i) => {
      if (!/[aeiou]/i.test(word[0])) {
        word = word.substring(1) + word[0];
      }
      return word + 'ma' + 'a'.repeat(i + 1);
    })
    .join(' ');
};
