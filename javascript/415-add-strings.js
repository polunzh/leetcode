/**
 * @param {string} num1
 * @param {string} num2
 * @return {string}
 */
var addStrings = function(num1, num2) {
    let n1 = 0;
    let n2 = 0;
    for(let i = num1.length - 1; i >= 0; i--){
        n1 += (num1[i].charCodeAt() - 48) * Math.pow(10, num1.length - i - 1);
    }

    for(let i = num2.length - 1; i >= 0; i--){
        n2 += (num2[i].charCodeAt() - 48) * Math.pow(10, num2.length - i - 1);
    }

    return String(n1 + n2);
};

console.log(addStrings('21', '22'));
