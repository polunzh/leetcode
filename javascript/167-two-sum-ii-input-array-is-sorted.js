const assert = require('assert');
/**
 * @param {number[]} numbers
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function(numbers, target) {
    let left = 0;
    let right = numbers.length - 1;

    while(left < right) {
        const val = numbers[left] + numbers[right];
        if (val === target) {
            return [left + 1, right + 1];
        } else if (val > target) {
            right--;
        } else {
            left++;
        }
    }
};

assert.deepEqual(twoSum([-1, 0], -1), [1, 2]);
assert.deepEqual(twoSum([2, 7, 10], 12), [1, 3]);
