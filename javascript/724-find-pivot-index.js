/**
 * @param {number[]} nums
 * @return {number}
 */
var pivotIndex = function(nums) {
  const sum = nums.reduce((acc, cur) => (acc += cur), 0);
  let leftSum = 0;
  for (let i = 0; i < nums.length; i++) {
    if (leftSum === sum - (leftSum + nums[i])) {
      return i;
    }
    leftSum += nums[i];
  }

  return -1;
};

console.log(pivotIndex([-1, -1, -1, 0, 1, 1]));
console.log(pivotIndex([-1,-1,0,1,1,0]));
