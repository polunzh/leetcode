const assert = require('assert');

/**
 * @param {number} n
 * @return {string}
 */
var convertToTitle = function(n) {
  const map = {
    '1': 'A',
    '2': 'B',
    '3': 'C',
    '4': 'D',
    '5': 'E',
    '6': 'F',
    '7': 'G',
    '8': 'H',
    '9': 'I',
    '10': 'J',
    '11': 'K',
    '12': 'L',
    '13': 'M',
    '14': 'N',
    '15': 'O',
    '16': 'P',
    '17': 'Q',
    '18': 'R',
    '19': 'S',
    '20': 'T',
    '21': 'U',
    '22': 'V',
    '23': 'W',
    '24': 'X',
    '25': 'Y',
    '26': 'Z',
  };

  let result = '';
  while (n > 0) {
    result = map[((n - 1) % 26) + 1] + result;
    n = Math.floor((n - 1) / 26);
  }

  return result;
};

assert.equal(convertToTitle(1), 'A');
assert.equal(convertToTitle(28), 'AB');
assert.equal(convertToTitle(701), 'ZY');
