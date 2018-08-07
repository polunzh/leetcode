/**
 * @param {number[]} nums
 * @return {number}
 */
var majorityElement = function(nums) {
    var min = nums.length / 2;
    var dict = {};
    for(var i = 0; i < nums.length; i++) {
        if(!dict[nums[i]]) {
            dict[nums[i]] = 1;
            if (min < 1) {
                return nums[i];
            }
        } else {
            dict[nums[i]]++;
            if (dict[nums[i]] > min) {
                return nums[i];
            }
        }
    }
};

console.log(majorityElement([1]));
console.log(majorityElement([3,2,3]));
console.log(majorityElement([2,2,1,1,1,2,2]));
