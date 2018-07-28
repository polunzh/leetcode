/** 牛顿迭代法
 * http://blog.punkid.org/2008/02/28/compute-the-square-root-via-newtons-iteration/
 * @param {number} x
 * @return {number}
 */
var mySqrt = function (x) {
    var N = 1000,
        i = 1,
        res = 1;

    for (; i <= N; i++) {
        res = (res + x / res) / 2;
    }

    return Math.floor(res);
};

console.log(mySqrt(3));