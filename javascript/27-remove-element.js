/**
 * @param {number[]} nums
 * @param {number} val
 * @return {number}
 */
var removeElement = function (nums, val) {
    if (nums === null || nums.length === 0) return 0;
    let len = nums.length,
        cur = 0;

    for (let i = 0; i < len; i++) {
        if (val !== nums[i]) {
            nums[cur++] = nums[i];
        }
    }

    return cur;
};

let a1 = [1, 2, 3, 1];
console.log(removeElement(a1, 1));
console.log(a1);
let a2 = [];
console.log(removeElement(a2, 1));
console.log(a2);
let a3 = [2];
console.log(removeElement(a3, 3));
console.log(a3);