const func = require('./index');

describe('simplify-path', () => {
  test('#1', () => {
    expect(func('/home/')).toEqual('/home');
  });

  test('#2', () => {
    expect(func('/../')).toEqual('/');
  });

  test('#3', () => {
    expect(func('/home//foo/')).toEqual('/home/foo');
  });

  test('#4', () => {
    expect(func('/a/./b/../../c/')).toEqual('/c');
  });
});
