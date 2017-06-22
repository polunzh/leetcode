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

let str1 = 'cdefewef';
let str2 = undefined;
let str3 = 'd';
console.log(longestCommonPrefix([str1, str2, str3]));