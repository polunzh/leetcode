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

var searchPos = function (s, target, start, end) {
    for (let i = start; i < end; i++) {
        if (s[i] === target) return i;
    }

    return -1;
};

var lengthOfLongestSubstring2 = function (s) {
    let start = 0,
        result = 0;

    for (let i = 0; i < s.length; i++) {
        let nStart = searchPos(s, s[i], start, i);
        if (nStart !== -1) {
            if (result < i - start) {
                result = i - start;
            }
            result = i - start > result ? i - start : result;
            start = nStart + 1;
        }
    }

    let len = s.length - start;
    return len > result ? len : result;
}

console.log(lengthOfLongestSubstring2('dvdf'));
console.log(lengthOfLongestSubstring2('aab'));
console.log(lengthOfLongestSubstring2(''));
console.log(lengthOfLongestSubstring2('pwwkew'));
console.log(lengthOfLongestSubstring2('bbbb'));
console.log(lengthOfLongestSubstring2('abcabcbb'));