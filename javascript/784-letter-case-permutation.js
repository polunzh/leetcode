/**
 * @param {string} S
 * @return {string[]}
 */
var letterCasePermutation = function(S) {
  const result = [];
  result.push(S);
  for (let i = 0; i < S.length; i++) {
    const len = result.length;
    if (S[i] >= 'a' && S[i] <= 'z') {
      let c = S[i].toUpperCase();
      for (let idx = 0; idx < len; idx++) {
        result.push(
          `${result[idx].substring(0, i)}${c.toUpperCase()}${result[
            idx
          ].substring(i + 1)}`
        );
      }
    } else if (S[i] >= 'A' && S[i] <= 'Z') {
      let c = S[i].toLowerCase();
      for (let idx = 0; idx < len; idx++) {
        result.push(
          `${result[idx].substring(0, i)}${c}${result[idx].substring(i + 1)}`
        );
      }
    }
  }

  return result;
};