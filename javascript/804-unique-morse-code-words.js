/**
 * @param {string[]} words
 * @return {number}
 */
var uniqueMorseRepresentations = function(words) {
  var MOSE = [
    '.-',
    '-...',
    '-.-.',
    '-..',
    '.',
    '..-.',
    '--.',
    '....',
    '..',
    '.---',
    '-.-',
    '.-..',
    '--',
    '-.',
    '---',
    '.--.',
    '--.-',
    '.-.',
    '...',
    '-',
    '..-',
    '...-',
    '.--',
    '-..-',
    '-.--',
    '--..',
  ];

  var result = [];

  for (var i = 0; i < words.length; i++) {
    var tmp = '';
    for (var j = 0; j < words[i].length; j++) {
      tmp += MOSE[words[i][j].charCodeAt() - 97];
    }

    result.push(tmp);
  }

  return new Set(result).size;
};

console.log(uniqueMorseRepresentations(['gin', 'zen', 'gig', 'msg']));
