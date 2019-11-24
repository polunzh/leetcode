/**
 * @param {string} text
 * @return {number}
 */
var maxNumberOfBalloons = function(text) {
  var balloonChars = { b: 0, a: 0, l: 0, o: 0, n: 0 };
  for (var i = 0; i < text.length; i++) {
    if (balloonChars[text[i]] !== undefined) {
      balloonChars[text[i]]++;
    }
  }

  balloonChars['l'] = Math.floor(balloonChars['l'] / 2);
  balloonChars['o'] = Math.floor(balloonChars['o'] / 2);
  var res = balloonChars['b'];
  Object.keys(balloonChars).forEach(key => {
    if (balloonChars[key] < res) {
      res = balloonChars[key];
    }
  });

  return res;
};
