import assert from "assert"

console.log("测试语法:: 变量声明");

function mes(i) {
    return `language/variable/2.js 测试用例[${i}]`
}

let a = 1
assert.equal(a, 1, mes(1))
let b = 1.2
assert.equal(b, 1.2, mes(2))
let c = "agc"
assert.equal(c, "agc", mes(3))
let d = [1, 2, 3]
assert.deepEqual(d, [1, 2, 3], mes(4))
assert.deepEqual([1, ...d, ...d, 2], [1, 1, 2, 3, 1, 2, 3, 2], mes(5))
let e
assert.equal(e, undefined, mes(6))
let f = null
assert.equal(f, null, mes(7))
let g = Infinity
assert.equal(g, Infinity, mes(8))
g = 123
assert.equal(g, 123, mes(9))