const assert = require('assert');
/**
 * initialize your data structure here.
 */
var MinStack = function() {
  this.minStack = [];
  this.stack = [];
};

/**
 * @param {number} x
 * @return {void}
 */
MinStack.prototype.push = function(x) {
  let i = 0;
  for (; i < this.minStack.length; i++) {
    if (this.minStack[i] >= x) {
      this.minStack.splice(i, -1, x);
      break;
    }
  }

  if (i === this.minStack.length) {
    this.minStack[i] = x;
  }

  this.stack.push(x);
};

/**
 * @return {void}
 */
MinStack.prototype.pop = function() {
  const element = this.stack.pop();
  if(element===this.minStack[0]){
    this.minStack.shift();
  }

  return element;
};

/**
 * @return {number}
 */
MinStack.prototype.top = function() {
  return this.stack[this.stack.length - 1];
};

/**
 * @return {number}
 */
MinStack.prototype.getMin = function() {
  return this.minStack[0];
};

MinStack.createNew = function() {
  return new MinStack();
};

/**
 * Your MinStack object will be instantiated and called as such:
 * var obj = Object.create(MinStack).createNew()
 * obj.push(x)
 * obj.pop()
 * var param_3 = obj.top()
 * var param_4 = obj.getMin()
 */

var obj = Object.create(MinStack).createNew();
obj.push(1);
obj.push(2);
assert(obj.top(), 2);
assert.equal(obj.getMin(), 1);
assert(obj.pop(), 2);
assert.equal(obj.getMin(), 1);
assert.equal(obj.top(), 1);