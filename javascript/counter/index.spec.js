const createCounter = require('./index');

describe('counter', () => {
  test('#1', () => {
    const counter = createCounter(10);
    expect(counter()).toBe(10);
    expect(counter()).toBe(11);
    expect(counter()).toBe(12);
  });
});
