/**
 * @param {number} n
 * @return {number}
 */
var subtractProductAndSum = function(n) {
  if (n === 0) {
    return 0;
  }

  var product = 1;
  var sum = 0;

  while (n > 0) {
    var tmp = n % 10;

    product *= tmp;
    sum += tmp;

    n = Math.floor(n / 10);
    console.log(tmp, n);
  }

  return product - sum;
};

console.log(subtractProductAndSum(234));
