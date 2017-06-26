/**
 * @param {string} s
 * @param {number} numRows
 * @return {string}
 */
var convert = function (s, numRows) {
    let len = s.length;
    let row = numRows;
    let col = parseInt(len / (2 * numRows - 2)) * (numRows - 1) + len % (2 * numRows - 2) % numRows + 1;
    let arr = new Array(row).fill(new Array(col).fill(0));
    let flag = true;

    

    console.log(arr);
};

convert('abcdefg', 3);