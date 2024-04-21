/**
 * @param {string} path
 * @return {string}
 */
const simplifyPath = function (path) {
  if (!path) {
    return '';
  }

  const strList = path.split('/');
  const stack = [];

  for (let i = 0; i < strList.length; i++) {
    if (strList[i] === '') {
      continue;
    }

    if (strList[i] === '.') {
      continue;
    }

    if (strList[i] === '..') {
      stack.pop();
      continue;
    }

    stack.push(strList[i]);
  }

  return `/${stack.join('/')}`;
};

module.exports = simplifyPath;
