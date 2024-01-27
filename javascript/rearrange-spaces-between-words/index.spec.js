const func = require('./index');

describe('rearrange-spaces-between-words', () => {
  test('#1', () => {
    expect(func('  this   is  a sentence ')).toBe('this   is   a   sentence');
  });

  test('#2', () => {
    expect(func('  hello')).toBe('hello  ');
  });
});
