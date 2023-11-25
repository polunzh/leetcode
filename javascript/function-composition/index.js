/**
 * @param {Function[]} functions
 * @return {Function}
 */
const compose = function (functions) {
  return function (x) {
    return functions.reduceRight((acc, func) => {
      return func(acc);
    }, x);
  };
};

module.exports = compose;
