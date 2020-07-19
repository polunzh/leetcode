// 使用”快慢指针“思想

/**
 * tool function
 * @param {number} n 
 */
const sum = (n) => {
  let result = 0;
  while (n > 9) {
    result += Math.pow(n % 10, 2);
    n = Math.floor(n / 10);
  }

  return result + Math.pow(n, 2);
};

/**
 * @param {number} n
 * @return {boolean}
 */
var isHappy = function (n) {
  let slow = n;
  let fast = n;
  do {
    slow = sum(slow);
    fast = sum(fast);
    fast = sum(fast);
  } while (slow !== fast);

  return slow === 1;
};

console.log(isHappy(2));
