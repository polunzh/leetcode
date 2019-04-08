/**
 * @param {number} N
 * @return {number}
 */
var fib = function(N) {
  if (N === 0) {
    return 0;
  }

  let counter = 2;
  let first = 0;
  let second = 1;
  while (counter++ <= N) {
    [first, second] = [second, first + second];
  }

  return second;
};

console.log(fib(1));
console.log(fib(2));
console.log(fib(3));
console.log(fib(4));
console.log(fib(5));
console.log(fib(10));
