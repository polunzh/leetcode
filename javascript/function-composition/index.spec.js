const compose = require('./index');

describe('compose', () => {
  test('#1', () => {
    const func = compose([(x) => x + 1, (x) => x * x, (x) => 2 * x]);
    expect(func(4)).toEqual(65);
  });

  test('#2', () => {
    const func = compose([x => 10 * x, x => 10 * x, x => 10 * x]);
    expect(func(1)).toEqual(1000);
  });

  test('#3', () => {
    const func = compose([]);
    expect(func(5)).toEqual(5);
  });
});
