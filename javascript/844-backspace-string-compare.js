/**
 * @param {string} S
 * @param {string} T
 * @return {boolean}
 */
var backspaceCompare = function(S, T) {
  const tmp1 = [];
  const tmp2 = [];
  for (let i = 0; i < S.length; i++) {
    if (S[i] === '#') {
      tmp1.pop();
    } else {
      tmp1.push(S[i]);
    }
  }

  for (let i = 0; i < T.length; i++) {
    if (T[i] === '#') {
      tmp2.pop();
    } else {
      tmp2.push(T[i]);
    }
  }

  return tmp1.join('') === tmp2.join('');
};

backspaceCompare('xywrrmp', 'xywrrm#p');
