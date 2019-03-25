/**
 * @param {number} num
 * @return {number}
 */
var addDigits = function(num) {
  while (num >= 10) {
    let t = 0;
    while (num >= 10) {
      t += num % 10;
      num = parseInt(num / 10, 10);
    }

    num = t + num;
  }

  return num;
};

/**
 * @param {number} num
 * @return {number}
 */
var addDigits2 = function(num) {
  while (num > 9) {
    num = (num + '').split('').reduce((acc, cur) => +acc + +cur, 0);
  }

  return num;
};

console.log(addDigits2(38));
