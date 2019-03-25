/**
 * @param {string} a
 * @param {string} b
 * @return {string}
 */
var addBinary = function(a, b) {
  let result = '';
  let t = 0;
  let aIndex = a.length - 1;
  let bIndex = b.length - 1;
  while (aIndex >= 0 || bIndex >= 0) {
    const t1 = Number(a[aIndex--]) || 0;
    const t2 = Number(b[bIndex--]) || 0;
    const tmp = t + t1 + t2;

    switch (tmp) {
      case 0:
        result = '0' + result;
        t = 0;
        break;
      case 1:
        result = '1' + result;
        t = 0;
        break;
      case 2:
        result = '0' + result;
        t = 1;
        break;
      case 3:
        result = '1' + result;
        t = 1;
        break;

      default:
        throw new Error('illegal');
    }
  }

  if (t === 1) {
    result = '1' + result;
  }

  return result;
};

console.log(addBinary('1', '1'));
