const fs = require('fs');
const path = require('path');

const genJavaScriptCode = (problemName) => {
  const p = path.join(__dirname, 'javascript', problemName);
  if (fs.existsSync(p)) {
    console.log(`FAIL: Problem ${problemName} already exists.`);
    return;
  }

  fs.mkdirSync(p, { recursive: true });
  fs.writeFileSync(path.join(p, 'index.js'), '');
  fs.writeFileSync(
    path.join(p, 'index.spec.js'),
    `
const func = require('./index');

describe('${problemName}', () => {
  test('#1', () => {});
});  
`
  );

  console.log(`SUCCESS: Problem ${path.relative(__dirname, p)} created.`);
};

/**
 *
 * @param {string} url Leetcode problem URL
 * @returns {string} Problem name
 */
const parseProblemName = (url) => {
  if (!url) {
    throw new Error('Please provide a leetcode URL.');
  }

  const obj = new URL(url);
  pathname = obj.pathname;
  return pathname.split('/')[2].toLowerCase();
};

const run = () => {
  if (process.argv.length < 3) {
    console.log('Please provide a leetcode URL.');
    return;
  }

  let [url, lang] = process.argv.slice(2);
  if (!lang) {
    lang = 'javascript';
  }

  const problemName = parseProblemName(url);
  if (lang === 'javascript') {
    genJavaScriptCode(problemName);
    return;
  }

  throw new Error(`Unsupported language: ${lang}`);
};

(() => {
  run();
})();
