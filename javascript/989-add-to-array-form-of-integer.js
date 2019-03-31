/**
 * @param {number[]} A
 * @param {number} K
 * @return {number[]}
 */
var addToArrayForm = function(A, K) {
  const kArray = [];
  while (K > 9) {
    kArray.unshift(K % 10);
    K = Math.floor(K / 10);
  }

  kArray.unshift(K);

  const result = [];

  const length = A.length > kArray.length ? A.length : kArray.length;
  let t = 0;
  for (let i = 0; i < length; i++) {
    const aIndex = A.length - i - 1;
    const kIndex = kArray.length - i - 1;
    const tmp =
      (aIndex > -1 ? A[aIndex] : 0) + (kIndex > -1 ? kArray[kIndex] : 0) + t;
    result.unshift(tmp % 10);
    if (tmp > 9) {
      t = 1;
    } else {
      t = 0;
    }
  }

  if (t === 1) {
    result.unshift(t);
  }

  return result;
};

console.log(addToArrayForm([9, 9, 9, 9, 9, 9, 9, 9, 9, 9], 1));
