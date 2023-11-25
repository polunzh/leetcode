/**
 * @param {*} obj
 * @param {*} classFunction
 * @return {boolean}
 */
const checkIfInstanceOf = function (obj, classFunction) {
  while (obj != null) {
    if (obj.constructor === classFunction) {
      return true;
    }

    obj = Object.getPrototypeOf(obj);
  }

  return false;
};

module.exports = checkIfInstanceOf;
