/**
 * @param {string} s
 * @return {number}
 */
var balancedStringSplit = function(s) {
  var stack = [];
  var rCount = 0;
  var lCount = 0;
  var count = 0;
  for (var i = 0; i < s.length; i++) {
    if (s[i] === 'R') {
      rCount++;
    } else {
      lCount++;
    }

    if(rCount===lCount){
      stack = [];
      count++;
    }
  }

  return count;
};

console.log(balancedStringSplit('RLRRLLRLRL'));
console.log(balancedStringSplit('RLLLLRRRLR'));
console.log(balancedStringSplit('LLLLRRRR'));
console.log(balancedStringSplit('RLRRRLLRLL'));
