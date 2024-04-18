import assert from "assert"

console.log("测试:: map");

function mes(i) {
    return `object/map/1.js 测试用例[${i}]`
}

let m = new Map()
m.set("x1", 1)
m.set("x2", 1)
m.set("x3", 1)
m.set("x4", 1)
assert.ok(m.size === 4, mes(1))
assert.deepEqual(m.keys(), ["x1", "x2", "x3", "x4"], mes(2))
assert.deepEqual(m.values(), [1, 1, 1, 1], mes(3))
