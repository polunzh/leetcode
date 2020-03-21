/**
 * @param {number[]} nums
 * @return {number[]}
 */
var decompressRLElist = function(nums) {
  const res = [];
  for (let i = 0; i < nums.length; i += 2) {
    res.splice(res.length, 0, ...new Array(nums[i]).fill(nums[i + 1]));
  }

  return res;
};

console.log(decompressRLElist([1, 1, 2, 3]));
