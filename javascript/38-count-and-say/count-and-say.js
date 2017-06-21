/**
 * @param {number} n
 * @return {string}
 */
var countAndSay = function (n) {
    if (n === 1) return '1';
    if (n === 2) return '11';
    if (n === 3) return '21';
    if (n === 4) return '1211';
    if (n === 5) return '111221';

    let preSerial = '1',
        temp,
        cur,
        count = 0;
    for (let i = 2; i <= n; i++) {
        temp = '';
        cur = preSerial[0];
        count = 1;
        for (let j = 0; j < preSerial.length; j++) {
            if (preSerial[j] !== preSerial[j + 1]) {
                temp += count + cur;
                cur = preSerial[j + 1];
                count = 1;
            } else {
                count++;
            }
        }

        preSerial = temp;
    }

    return preSerial;
};

console.log(countAndSay(1));
console.log(countAndSay(2));
console.log(countAndSay(3));
console.log(countAndSay(5));
console.log(countAndSay(10));