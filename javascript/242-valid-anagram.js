/**
 * @param {string} s
 * @param {string} t
 * @return {boolean}
 */
var isAnagram = function(s, t) {
  if (s.length !== t.length) {
    return false;
  }

  var sArr = s.split('');
  var tArr = t.split('');

  for (let i = 0; i < sArr.length; i++) {
    var flag = false;
    for (let j = i; j < tArr.length; j++) {
      if (sArr[i] === tArr[j]) {
        flag = true;
        if (i !== j) {
          var tmp = tArr[i];
          tArr[i] = tArr[j];
          tArr[j] = tmp;
        }

        break;
      }
    }

    if (!flag) {
      return false;
    }
  }

  return true;
};

console.log(isAnagram('a', ''));
console.log(isAnagram('a', 'a'));
console.log(isAnagram('a', 'b'));;
console.log(isAnagram('anagram', 'nagaram'));
console.log(isAnagram('rat', 'car'));
console.log(isAnagram('aacc', 'ccac'));
