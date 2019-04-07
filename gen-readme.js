const fs = require('fs');
const GITHUB_URL = 'https://github.com/polunzh/leetcode/blob/master/';
const LEETCODE_URL = 'https://leetcode.com/problems/';

const jsSolutions = fs.readdirSync('javascript');
const pySolutions = fs.readdirSync('python');

const fileTypeExtensionMap = new Map();
fileTypeExtensionMap.set('javascript', 'js');
fileTypeExtensionMap.set('python', 'py');

/**
 *
 * @param {number} order
 * @param {string} name
 * @param {string[]} fileTypes
 */
const generateTableBodyRows = (order, name, fileTypes) => {
  const solutions = fileTypes
    .reduce((acc, cur) => {
      acc.push(
        `[${cur}](${GITHUB_URL}${cur.toLowerCase()}/${order}-${name}.${fileTypeExtensionMap.get(
          cur
        )})`
      );

      return acc;
    }, [])
    .join(',');

  return `| ${order} | [${name}](${LEETCODE_URL}${name})  | ${solutions} |`;
};

/**
 *
 * @param {Map<number,string>} map
 * @param {string[]} solutions
 * @param {string} type
 */
const generateSolutionsMap = (map, solutions, type) => {
  solutions.forEach(item => {
    const tmpIndex = item.indexOf('-');
    const order = Number(item.substring(0, tmpIndex));
    const name = item.substring(tmpIndex + 1, item.lastIndexOf('.'));
    if (map.get(order)) {
      map.get(order).solutions.push(type);
    } else {
      map.set(order, { name, solutions: [type] });
    }
  });
};

let map = new Map();
generateSolutionsMap(map, jsSolutions, 'javascript');
generateSolutionsMap(map, pySolutions, 'python');

map = new Map(
  [...map.entries()].sort((x, y) => {
    x = x[0];
    y = y[0];

    return x - y;
  })
);

const tableBody = [];
for (let [order, value] of map) {
  tableBody.push(generateTableBodyRows(order, value.name, value.solutions));
}

const tableHeader =
  '# LeetCode\n\n | # | Title  | Solutions | \n |---|---|---| \n';

const content = `${tableHeader}${tableBody.join('\n')}`;
fs.writeFileSync('README.md', content);
console.log('README.md file generate success!');
