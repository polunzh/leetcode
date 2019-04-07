/**
 * @param {number} num
 * @return {string}
 */
var convertToBase7 = function(num) {
  const flag = num < 0 ? -1 : 1;
  num = flag * num;

  let result = '';
  while (num > 0) {
    result = (num % 7) + result;
    num = Math.floor(num / 7);
  }

  return result * flag + '';
};

console.log(convertToBase7(-7));
