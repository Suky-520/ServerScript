import assert from "assert"

console.log("测试语法 while");

let a = 1;
while (true) {
    a++;
    if (a > 10) {
        break;
    }
}
assert.equal(a, 11, mes(1));

a = 1;
while (a < 10) {
    a++;
}
assert.equal(a, 10, mes(2));

function mes(i) {
    return `language/for/2.js 测试用例[${i}]`;
}