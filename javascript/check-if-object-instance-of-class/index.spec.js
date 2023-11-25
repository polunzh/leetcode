const checkIfInstanceOf = require('./index.js');

describe('2618-check-if-object-instance-of-class', () => {
  test('#1', () => {
    expect(checkIfInstanceOf(new Date(), Date)).toEqual(true);
  });

  test('#2', () => {
    class Animal {}
    class Dog extends Animal {}

    expect(checkIfInstanceOf(new Dog(), Animal)).toEqual(true);
  });

  test('#3', () => {
    expect(checkIfInstanceOf(Date, Date)).toEqual(false);
  });

  test('#4', () => {
    expect(checkIfInstanceOf(5, Number)).toEqual(true);
  });
});
