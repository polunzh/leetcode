/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var search = function(nums, target) {
  let low = 0;
  let hi = nums.length - 1;
  while (low <= hi) {
    let mid = Math.floor(low + (hi - low) / 2);
    if (nums[mid] === target) {
      return mid;
    } else if (nums[mid] > target) {
      hi = mid - 1;
    } else {
      low = mid + 1;
    }
  }

  return -1;
};

console.log(search([-1, 0, 1, 2], 1));
