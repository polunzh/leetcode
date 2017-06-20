/**
 * @param {number} n
 * @return {number}
 */
var climbStairs = function (n) {
    if (n <= 0) return 0;

    var a = 0,
        b = 0,
        x = parseInt(n / 2),
        res = 0, i;

    while ((a >= 0 && a <= n)
        && (b >= 0 && b <= x)) {
        if ((n - a) % 2 === 0) {
            b = (n - a) / 2;

            // 求排列组合-组合: C(b)/(n-b)
            if (a !== 0 && b !== 0) {
                var t0 = n - b,
                    t1 = 1,
                    t2 = 1;
                for (i = 0; i < b; i++) {
                    t1 *= (n - b - i);
                    t2 *= (b - i);
                }

                res += t1 / t2;
            } else { res++; }
        }

        a++;
    }

    return res;
};

console.log(climbStairs(4));