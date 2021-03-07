const assert = require('assert').strict;

/**
 * @param {string} name
 * @param {string} typed
 * @return {boolean}
 */
var isLongPressedName = function (name, typed) {
  let i = 0;
  let j = 0;

  for (; i < name.length && j < typed.length; ) {
    if (name[i] != typed[j]) {
      return false;
    }

    let tmp1 = 1;
    let tmp2 = 1;
    for (; name[i + tmp1] === name[i]; tmp1++) {}
    for (; typed[j + tmp2] === typed[j]; tmp2++) {}

    if (tmp1 > tmp2) {
      return false;
    }

    i += tmp1;
    j += tmp2;
  }

  if (typeof name[i] !== 'undefined' || typeof typed[j] !== 'undefined') {
    return false;
  }

  return true;
};

assert.equal(isLongPressedName('alex', 'aaleex'), true);
assert.equal(isLongPressedName('saeed', 'ssaaedd'), false);
assert.equal(isLongPressedName('leelee', 'lleeelee'), true);
assert.equal(isLongPressedName('laiden', 'laiden'), true);
assert.equal(isLongPressedName('xnhtq', 'xhhttqq'), false);
assert.equal(isLongPressedName('alex', 'aaleexa'), false);
