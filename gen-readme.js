const fs = require('fs');
const githubUrl = 'https://github.com/polunzh/leetcode/blob/master/';
const leetcodeUrl = 'https://leetcode.com/problems/';
const spliter = '<!--table-->';

const jsSolutions = fs.readdirSync('javascript');
const pySolutions = fs.readdirSync('python');
const strMap = new Map();
let name, idx, temp;

function genStr(s, typeName, fileExt) {
    idx = s.indexOf('-');
    name = s.substring(idx + 1);
    return `| ${s.substring(0, idx)} | [${name}](${leetcodeUrl}${name})  |  [${typeName}](${githubUrl}${typeName.toLowerCase()}/${s}/${name}.${fileExt})|`;
}

jsSolutions.forEach((item) => {
    strMap.set(item, genStr(item, 'JavaScript', 'js'));
});

console.log(jsSolutions);

pySolutions.forEach((item) => {
    if (strMap.has(item)) {
        temp = strMap.get(item);
        temp = temp.substring(0, temp.lastIndexOf('|'));
        temp += ` , [Python](${githubUrl}python/${item}/${item.substring(item.indexOf('-') + 1)}.py) |`;
        strMap.set(item, temp);
    } else {
        strMap.set(item, genStr(item, 'Python', 'py'));
    }
});

const newMap = new Map([...strMap.entries()].sort((a, b) => {
    a = a[0];
    b = b[0];
    let aLen = Number(a.substring(0, a.indexOf('-'))),
        bLen = Number(b.substring(0, b.indexOf('-')));
    if (aLen < bLen) return -1;
    else if (aLen > bLen) return 1;
    else return 0;
}));

/*
let oriContent = fs.readFileSync('README.md');
let contentArray = oriContent.toString().split(spliter);
contentArray[1] = spliter + `\n\n | # | Title  | Solution | \n |---|---|---| \n` + Array.from(newMap.values()).join('\n');
fs.writeFile('README.md', contentArray.join(' '));
console.log('README.md file generate success!');
*/