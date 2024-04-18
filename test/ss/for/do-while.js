import assert from "assert"

console.log("测试语法 do-while");

function mes(i) {
    return `language/for/3.js 测试用例[${i}]`
}

let a = 1;
do {
    a++;
} while (a < 10)

assert.equal(a, 10, mes(1))

a = 1
do {
    a++
    if (a > 10) {
        break
    }
} while (true)

assert.equal(a, 11, mes(2))
