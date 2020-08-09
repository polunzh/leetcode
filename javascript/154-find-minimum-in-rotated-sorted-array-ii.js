const assert = require('assert');
/**
 * @param {number[]} nums
 * @return {number}
 */
const findMin = function (nums) {
  if (nums.length === 1) {
    return nums[0];
  }

  let firstIndex = 0;
  let secondIndex = nums.length - 1;
  let midIndex = firstIndex;

  while (nums[firstIndex] >= nums[secondIndex]) {
    if (secondIndex - firstIndex === 1) {
      midIndex = secondIndex;
      break;
    }

    if (
      nums[firstIndex] === nums[secondIndex] &&
      nums[midIndex] == nums[firstIndex]
    ) {
      let res = nums[firstIndex];
      for (let i = firstIndex + 1; i < secondIndex; i++) {
        if (nums[i] < res) {
          res = nums[i];
        }
      }

      return res;
    }

    midIndex = firstIndex + Math.floor((secondIndex - firstIndex) / 2);
    if (nums[firstIndex] <= nums[midIndex]) {
      firstIndex = midIndex;
    } else {
      secondIndex = midIndex;
    }
  }

  return nums[midIndex];
};

// assert.equal(findMin([1]), 1);
// assert.equal(findMin([3, 1, 3]), 1);
// assert.equal(findMin([1, 3, 5]), 1);
// assert.equal(findMin([2, 2, 2, 0, 1]), 0);
assert.equal(findMin([10, 1, 10, 10, 10]), 1);
