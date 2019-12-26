/**
 * @param {number[]} stones
 * @return {number}
 */
var lastStoneWeight = function(stones) {
  stones.sort((x, y) => y - x);
  while (stones.length > 1) {
    const tmp = Math.abs(stones.shift() - stones.shift());
    let i = 0;
    while (stones[i] >= tmp) {
      i++;
    }

    stones.splice(i, 0, tmp);
  }

  if (stones.length === 0) {
    return 0;
  }
  return stones[0];
};

lastStoneWeight([2, 7, 4, 1, 8, 1]);
