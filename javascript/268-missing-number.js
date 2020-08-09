const assert = require('assert');
/**
 * @param {number[]} nums
 * @return {number}
 */
const missingNumber = function (nums) {
  let ifNoneSum = nums.length;
  let sum = 0;
  for (let i = 0; i < nums.length; i++) {
    ifNoneSum += i;
    sum += nums[i];
  }

  return ifNoneSum - sum;
};

assert.equal(missingNumber([3, 0, 1]), 2);
assert.equal(missingNumber([9, 6, 4, 2, 3, 5, 7, 0, 1]), 8);
