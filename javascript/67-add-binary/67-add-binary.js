const assert = require('assert');
/** 做的漏洞百出
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function (a, b) {
    var arr1, arr2;

    if (a.length < b.length) {
        arr1 = a;
        arr2 = b;
    } else {
        arr1 = b;
        arr2 = a;
    }

    var len1 = arr1.length,
        len2 = arr2.length,
        e = 0,
        s = Math.abs(len1 - len2),
        newArr = [], t;

    for (var i = len1 - 1; i >= 0; i--) {
        t = Number(arr1[i]) + Number(arr2[i + s]) + e;
        if (t === 3) {
            newArr.unshift(1);
            e = 1;
        } else if (t === 2) {
            newArr.unshift(0);
            e = 1;
        } else if (t === 1) {
            newArr.unshift(1);
            e = 0;
        } else {
            newArr.unshift(0);
            e = 0;
        }

    }

    for (var i = len2 - len1 - 1; i >= 0; i--) {
        t = Number(arr2[i]) + e;
        if (t === 2) {
            newArr.unshift(0);
            e = 1;
        } else if (t === 1) {
            newArr.unshift(1);
            e = 0;
        } else {
            newArr.unshift(0);
            e = 0;
        }
    }

    if (e === 1) {
        newArr.unshift(1);
    }

    return newArr.join('');
};
assert.equal(addBinary('100', '110010'), '110110', 'FAIL');
assert.equal(addBinary('0', '1'), '1', 'FAIL');
assert.equal(addBinary('1', '1'), '10', 'FAIL');
assert.equal(addBinary('11', '1'), '100', 'FAIL');
assert.equal(addBinary('1', '11'), '100', 'FAIL');
// console.log(addBinary('1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111', '111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111'));