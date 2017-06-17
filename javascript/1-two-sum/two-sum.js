/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number[]}
 */
var twoSum = function (nums, target) {
    map = {}
    var x;
    for (var i = 0; i < nums.length; i++) {
        x = map[target - nums[i]];
        if (x !== undefined) {
            return [x, i];
        }

        map[nums[i]] = i;
    }

    throw new Error('Not found');
};

console.log(twoSum([3, 2, 4], 6))