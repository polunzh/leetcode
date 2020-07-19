// Symbol       Value
// I             1
// V             5
// X             10
// L             50
// C             100
// D             500
// M             1000

// I can be placed before V (5) and X (10) to make 4 and 9.
// X can be placed before L (50) and C (100) to make 40 and 90.
// C can be placed before D (500) and M (1000) to make 400 and 900.

// MCMXCIV

/**
 * @param {string} s
 * @return {number}
 */
const romanToInt = function (s) {
  let result = 0;
  let last = null;
  const map = { I: 1, V: 5, X: 10, L: 50, C: 100, D: 500, M: 1000 };
  const rule = {
    I: ['V', 'X'],
    X: ['L', 'C'],
    C: ['D', 'M'],
  };

  for (let i = 0; i < s.length; i++) {
    result += map[s[i]];
    if (last !== null && rule[last] && rule[last].includes(s[i])) {
      result -= 2 * map[last];
    }

    last = s[i];
  }

  return result;
};

console.log(romanToInt('MCMXCIV'));
console.log(romanToInt('IX'));
