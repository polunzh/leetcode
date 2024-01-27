/**
 * @param {string} text
 * @return {string}
 */
var reorderSpaces = function (text) {
  const arr = text.split(/\s+/).filter(Boolean);
  const spaceCount =
    text.length - arr.reduce((acc, cur) => acc + cur.length, 0);
  const averageSpaceCount =
    arr.length === 1 ? spaceCount : Math.floor(spaceCount / (arr.length - 1));
  const restSpaceCount = spaceCount - averageSpaceCount * (arr.length - 1);

  const res = arr.join(' '.repeat(averageSpaceCount));
  if (restSpaceCount === 0) {
    return res;
  }

  return `${res}${' '.repeat(restSpaceCount)}`;
};

module.exports = reorderSpaces;
