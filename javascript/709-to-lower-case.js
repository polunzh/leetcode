/**
 * @param {string} str
 * @return {string}
 */
var toLowerCase = function(str) {
  var res = '';
  for (var i = 0; i < str.length; i++) {
    if (str[i].charCodeAt() >= 65 && str[i].charCodeAt() <= 90) {
      res += String.fromCharCode(97 + str[i].charCodeAt() - 65);
    } else {
      res += str[i];
    }
  }

  return res;
};

console.log(toLowerCase('abc'));
