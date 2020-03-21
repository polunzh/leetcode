/**
 * @param {number} num
 * @return {number}
 */
var numberOfSteps = function(num) {
  let steps = 0;
  while (num > 0) {
    steps++;
    if (num % 2 !== 0) {
      num--;
      continue;
    }

    if (num === 0) {
      break;
    }

    num /= 2;
  }

  return steps;
};
