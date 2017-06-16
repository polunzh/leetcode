/**
 * @param {number} x
 * @return {number}
 */
var reverse = function (x) {
    var flag = x < 0 ? -1 : 1,
        res = 0,
        n;

    while(x) {
        n = x % 10;
        res = res * 10 + x % 10;
        x = parseInt(x / 10);
    }

    if (res < -2147483648 || res > 2147483647) return 0;

    return res;
};

console.log(reverse(0), -321);