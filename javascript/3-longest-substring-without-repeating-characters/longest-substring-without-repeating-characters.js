/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLongestSubstring = function (s) {
    if (s.length === 0) return 0;
    let cur = 0,
        result = [],
        arr = [],
        start = 0;

    for (let i = 0; i < s.length;) {
        let idx = arr.indexOf(s[i]);
        if (idx !== -1) {
            if (result.length < arr.length) {
                result = Array.from(arr);
            }

            i = (i - arr.length) + 1;
            arr.push(s[i]);
            arr = [];
        } else {
            arr.push(s[i]);
            i++;
        }
    }

    if (result.length < arr.length) {
        result = Array.from(arr);
    }
    return result.length;
};

console.log(lengthOfLongestSubstring('dvdf'));
console.log(lengthOfLongestSubstring('aab'));
console.log(lengthOfLongestSubstring(''));
console.log(lengthOfLongestSubstring('pwwkew'));
console.log(lengthOfLongestSubstring('bbbb'));
console.log(lengthOfLongestSubstring('abcabcbb'));