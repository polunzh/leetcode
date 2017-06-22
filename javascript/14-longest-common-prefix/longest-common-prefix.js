/**
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix = function (strs) {
    if (!strs || strs.length === 0) return '';

    strs.sort();
    let prefix = strs[0],
        lastStr = strs[strs.length - 1],
        len = prefix.length;

    for (let i = 0; i < len; i++) {
        if (prefix[i] !== lastStr[i]) return prefix.substring(0, i);
    }

    return prefix;
};

/** 别人的解法
 * @param {string[]} strs
 * @return {string}
 */
var longestCommonPrefix2 = function (strs) {
    if (strs.length === 0) {
        return "";
    }

    let prefix = strs[0];

    for (let i = 1; i < strs.length; i++) {
        while (prefix.length >= 0) {
            if (strs[i].indexOf(prefix) !== 0) {
                prefix = prefix.substring(0, prefix.length - 1);
            } else {
                break;
            }
        }

        if (prefix.length === 0) {
            break;
        }
    }

    return prefix;
};

let str1 = 'abcdefewef';
let str2 = 'abce';
let str3 = 'abcd';
console.log(longestCommonPrefix2([str1, str2, str3]));