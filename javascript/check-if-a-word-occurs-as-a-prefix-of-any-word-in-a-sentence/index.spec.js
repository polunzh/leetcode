const func = require('./index');

describe('check-if-a-word-occurs-as-a-prefix-of-any-word-in-a-sentence', () => {
  test('#1', () => {
    expect(func('i love eating burger', 'burg')).toEqual(4);
  });

  test('#2', () => {
    expect(func('this problem is an easy problem', 'pro')).toEqual(2);
  });

  test('#3', () => {
    expect(func('i am tired', 'you')).toEqual(-1);
  });
});
