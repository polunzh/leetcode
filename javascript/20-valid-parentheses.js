/**
 * @param {string} s
 * @return {boolean}
 */
var isValid = function (s) {
    var stack = [],
        len = s.length,
        map = {
            ')': '(',
            ']': '[',
            '}': '{',
        };
    for (var i = 0; i < len; i++) {
        if (s[i] === '(' || s[i] === '[' || s[i] === '{') {
            stack.push(s[i]);
        } else {
            if (map[s[i]] !== stack[stack.length - 1]) return false;

            stack.pop();
        }
    }

    return stack.length <= 0;
};

console.log(isValid('()[]{}'));
console.log(isValid(']'));
console.log(isValid('()[{]}'));