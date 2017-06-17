/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome = function (x) {
    if (x < 0) return false;
    return String(x).split('').reverse().join('') === String(x);
};

/**
 * @param {number} x
 * @return {boolean}
 */
var isPalindrome2 = function (x) {
    if (x < 0) return false;
    if (x > 0 && x % 10 === 0) return false;

    n = x;
    m = 0;
    while (m < n) {
        m = m * 10 + n % 10;
        n = parseInt(n / 10);
    }

    return m === n || parseInt(m / 10) == n;
};

console.log(isPalindrome2(12321));