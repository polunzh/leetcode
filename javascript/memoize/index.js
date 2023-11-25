/**
 * @param {Function} fn
 * @return {Function}
 */
function memoize(fn) {
  const cache = {};

  return function (...args) {
    const key = JSON.stringify(args);

    if (key in cache) {
      return cache[key];
    }

    const value = fn(...args);
    cache[key] = value;

    return value;
  };
}

module.exports = memoize;
