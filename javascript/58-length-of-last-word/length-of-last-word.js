/**
 * @param {string} s
 * @return {number}
 */
var lengthOfLastWord = function (s) {
    s = s.trim();
    var idx = s.lastIndexOf(' ');
    if (idx === s.length - 1) return 0;

    return s.length - idx - 1;
};

console.log(lengthOfLastWord(" i   "));
console.log(lengthOfLastWord("Hello   World"));