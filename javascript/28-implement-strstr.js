/**
 * @param {string} haystack
 * @param {string} needle
 * @return {number}
 */
var strStr = function (haystack, needle) {
    if (haystack.length < needle.length) return -1;
    if (haystack === needle) return 0;

    let len = haystack.length,
        i, j;
    for (i = 0; i < haystack.length - needle.length + 1; i++) {
        for (j = 0; j < needle.length; j++) {
            if (needle[j] !== haystack[i + j]) break;
        }

        if (j === needle.length) return i;
    }

    return -1;
};
console.log(strStr('zhangz', ''));
console.log(strStr('zhangzhenqiang', 'ia'));