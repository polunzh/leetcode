/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function (s, numRows) {
    if (s.length <= numRows || numRows === 1) return s;
    let len = s.length;
    let row = numRows;
    let col = parseInt(len / (2 * numRows - 2)) * (numRows - 1) + 1;
    let y = len % (2 * numRows - 2);
    if (y > numRows) {
        col += y - numRows;
    }

    let arr = [];
    for (let i = 0; i < row; i++) {
        arr[i] = new Array(col).fill(0);
    }

    let flag = true;

    for (let i = 0; i < len; i++) {
        let t = (i + 1) % (2 * numRows - 2);
        let r;
        if (t === 0) r = 1;
        else if (t <= numRows) r = t - 1;
        else r = 2 * numRows - t - 1;

        let c = parseInt((i + 1) / (2 * numRows - 2)) * (numRows - 1);
        let y = (i + 1) % (2 * numRows - 2);
        if (y === 0) c--;
        else if (y > numRows) c += y - numRows;
        arr[r][c] = s[i];
    }

    let res = '';
    arr.forEach((a) => {
        a.forEach((b) => {
            if (b !== 0) res += b;
        });
    });
    return res;
};

console.log(convert('ABCDE', 4));