/* 统计排序思路 */

/**
 * @param {number[]} nums
 * @return {number[]}
 */
var smallerNumbersThanCurrent = function(nums) {
  const count = new Array(101).fill(0);
  const res = new Array(nums.length).fill(0);

  for (let i = 0; i < nums.length; i++) {
    count[nums[i]]++;
  }

  for (let i = 1; i < count.length; i++) {
    count[i] += count[i - 1];
  }

  for (let i = 0; i < nums.length; i++) {
    if (nums[i] === 0) {
      res[i] = 0;
    } else {
      res[i] = count[nums[i] - 1];
    }
  }

  console.log(res);
  return res;
};

smallerNumbersThanCurrent([0, 0, 0, 0, 0]);
