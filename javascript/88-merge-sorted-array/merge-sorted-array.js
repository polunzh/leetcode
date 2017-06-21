/** 第次一做的十分垃圾，失败！
 *  虽然实现了，但是需要优化
 * Given two sorted integer arrays nums1 and nums2, merge nums2 into nums1 as one sorted array.
 * Note:
 * You may assume that nums1 has enough space (size that is greater or equal to m + n) to hold additional elements from nums2.
 * The number of elements initialized in nums1 and nums2 are m and n respectively.
 * 
*/

var searchInsert = function (nums, left, target) {
    let len = nums.length,
        right = len - 1,
        mid;

    if (nums[0] >= target) return 0;

    if (nums[len - 1] < target) return len;

    if (nums[len - 1] == target) return len - 1;

    while (left <= right) {
        mid = left + parseInt((right - left) / 2);
        if (nums[mid] === target) {
            return mid + 1;
        } else if (nums[mid] < target) {
            left = mid + 1;
        } else {
            right = mid - 1;
        }
    }

    return nums[mid] < target ? mid + 1 : mid;
};

/**
 * @param {number[]} nums1
 * @param {number} m
 * @param {number[]} nums2
 * @param {number} n
 * @return {void} Do not return anything, modify nums1 in-place instead.
 */
var merge = function (nums1, m, nums2, n) {
    let pos, start = 0;


    while (m < nums1.length) {
        nums1.pop();
    }

    if (m === 0) {
        nums2.forEach((x) => nums1.push(x));
        return;
    }

    for (let i = 0; i < n; i++) {
        pos = searchInsert(nums1, start, nums2[i]);
        start = pos + 1;
        nums1.splice(pos, 0, nums2[i]);
    }
};

let nums1 = [1, 2];
let nums2 = [4, 6];
merge(nums1, nums1.length, nums2, nums2.length);
console.log(nums1);
nums1 = [1, 2, 6];
nums2 = [4, 6];
merge(nums1, nums1.length, nums2, nums2.length);
console.log(nums1);
nums1 = [10, 12, 16];
nums2 = [4, 6];
merge(nums1, nums1.length, nums2, nums2.length);
console.log(nums1);
nums1 = [0];
nums2 = [1];
merge(nums1, 0, nums2, nums2.length);
console.log(nums1);
nums1 = [1, 0];
nums2 = [2];
merge(nums1, 1, nums2, nums2.length);
console.log(nums1);
nums1 = [1, 2, 3, 0, 0, 0];
nums2 = [2, 5, 6];
merge(nums1, 3, nums2, nums2.length);
console.log(nums1);
nums1 = [-1, 0, 1, 1, 0, 0, 0, 0, 0];
nums2 = [-1, 0, 2, 2, 3];
merge(nums1, 4, nums2, nums2.length);
console.log(nums1);