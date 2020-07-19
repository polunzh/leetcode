/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
const search = function (nums, target) {
  let low = 0;
  let high = nums.length - 1;

  while (low < high) {
    let mid = Math.floor((low + high) / 2);
    if (nums[mid] > nums[high]) {
      low = mid + 1;
    } else {
      high = mid;
    }
  }

  let rotate = low;
  low = 0;
  high = nums.length - 1;
  while (low <= high) {
    let mid = Math.floor((low + high) / 2);
    let realMid = (mid + rotate) % nums.length;
    if (nums[realMid] === target) {
      return realMid;
    } else if (nums[realMid] < target) {
      low = mid + 1;
    } else {
      high = mid - 1;
    }
  }

  return -1;
};

console.log(search([2, 3, 4, -1, 0, 1], 1));
