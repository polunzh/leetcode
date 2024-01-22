/**
 * @param {string[]} words
 * @param {character} separator
 * @return {string[]}
 */
 var splitWordsBySeparator = function(words, separator) {
  var result = [];
  words.forEach(w => {
      w.split(separator).forEach(x => {
          if(x.trim()) {
              result.push(x);
          }
      });
  });

  return result;
};