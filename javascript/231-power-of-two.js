const assert = require('assert');
/**
 * @param {number} n
 * @return {boolean}
 */
var isPowerOfTwo = function(n) {
    if (n <= 0) {
        return false;
    }

    if (n === 1 ) {
        return true;
    }

    while (n % 2 === 0) {
        n = n / 2;
    }

    return n === 1;
};

assert.equal(isPowerOfTwo(0), false);
assert.equal(isPowerOfTwo(1), true);
assert.equal(isPowerOfTwo(4), true);
