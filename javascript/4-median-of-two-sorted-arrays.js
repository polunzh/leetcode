/**
 * @param {number[]} nums1
 * @param {number[]} nums2
 * @return {number}
 */
var findMedianSortedArrays = function (nums1, nums2) {
  while (nums1.length + nums2.length > 2) {
    if (nums1.length > 0 && nums2.length > 0) {
      if (nums1[0] < nums2[0]) {
        nums1.shift();
      } else {
        nums2.shift();
      }
    } else if (nums1.length > 0) {
      nums1.shift();
    } else {
      nums2.shift();
    }

    if (nums1.length > 0 && nums2.length > 0) {
      if (nums1[nums1.length - 1] > nums2[nums2.length - 1]) {
        nums1.pop();
      } else {
        nums2.pop();
      }
    } else if (nums1.length > 0) {
      nums1.pop();
    } else {
      nums2.pop();
    }
  }

  console.log(nums1, nums2);
  const res = [...nums1, ...nums2];
  return res.length === 2 ? (res[0] + res[1]) / 2 : res[0];
};

console.log(findMedianSortedArrays([3], [-2, -1]));
