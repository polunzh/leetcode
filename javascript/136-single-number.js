const assert = require('assert');

/**
 * @param {number[]} nums
 * @return {number}
 */
const singleNumber = function (nums) {
  const map = new Map();
  for (let i = 0; i < nums.length; i++) {
    if (map.has(nums[i])) {
      if (map.get(nums[i]) === 1) {
        map.delete(nums[i]);
      } else {
        map.set(nums[i], 1);
      }
    } else {
      map.set(nums[i], 1);
    }
  }

  return [...map.keys()][0];
};

assert.equal(singleNumber([2, 2, 1]), 1);
assert.equal(singleNumber([4, 1, 2, 1, 2]), 4);
