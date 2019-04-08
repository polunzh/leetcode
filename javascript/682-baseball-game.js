/**
 * @param {string[]} ops
 * @return {number}
 */
var calPoints = function(ops) {
  const stack = [];
  let sum = 0;
  for (let i = 0; i < ops.length; i++) {
    let t = 0;
    if (ops[i] === 'C') {
      t = -1 * stack.pop();
    } else if (ops[i] === 'D') {
      t = stack[stack.length - 1] * 2;
      stack.push(t);
    } else if (ops[i] === '+') {
      t = stack[stack.length - 1] + stack[stack.length - 2];
      stack.push(t);
    } else {
      t = Number(ops[i]);
      stack.push(t);
    }
    sum += t;
  }

  return sum;
};

console.log(calPoints(['5', '-2', '4', 'C', 'D', '9', '+', '+']));
