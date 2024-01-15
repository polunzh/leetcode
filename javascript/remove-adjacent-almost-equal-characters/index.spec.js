
const func = require('./index');

describe('remove-adjacent-almost-equal-characters', () => {
  test('#1', () => {
    expect(func('aaaaa')).toEqual(2)
  });

  test('#2', () => {
    expect(func('abddez')).toEqual(2)
  });

  test('#3', () => {
    expect(func('zyxyxyz')).toEqual(3)
  });

  test('#4', () => {
    expect(func('acba')).toEqual(1)
  });

  test('#5', () => {
    expect(func('aaaac')).toEqual(2)
  });

  test('#6', () => {
    expect(func('aacca')).toEqual(2)
  });
});  
