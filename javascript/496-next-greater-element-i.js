/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number[]}
 */
var nextGreaterElement = function(nums1, nums2) {
  const result = [];
  const map = [];
  const stack = [];
  for (let i = 0; i < nums2.length; i++) {
    while (stack.length && stack[stack.length - 1] < nums2[i]) {
      map[stack.pop()] = nums2[i];
    }

    stack.push(nums2[i]);
  }

  for (let i = 0; i < nums1.length; i++) {
    result.push(map[nums1[i]] || -1);
  }

  return result;
};

const assert = require('assert');

assert.deepEqual([-1, 3, -1], nextGreaterElement([4, 1, 2], [1, 3, 4, 2]));
assert.deepEqual([3, -1], nextGreaterElement([2, 4], [1, 2, 3, 4]));
