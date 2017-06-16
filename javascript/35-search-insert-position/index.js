/**
 * @param {number[]} nums
 * @param {number} target
 * @return {number}
 */
var searchInsert = function (nums, target) {

    var left = 0,
        len = nums.length,
        right = len - 1,
        mid;

    if (nums[0] >= target) return 0;

    if (nums[len - 1] < target) return len;

    if (nums[len - 1] == target) return len - 1;

    while (left <= right) {
        mid = left + parseInt((right - left) / 2);
        if (nums[mid] === target) {
            return mid;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    return nums[mid] < target ? mid + 1 : mid;
};