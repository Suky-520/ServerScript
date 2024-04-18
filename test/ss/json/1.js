import assert from "assert"

console.log("测试:: json序列化");

function mes(i) {
    return `object/json/1.js 测试用例[${i}]`
}

let a = 1
assert.ok(JSON.stringify(a) === "1", mes(1))
assert.ok(JSON.stringify(a) == "1", mes(2))
assert.ok(JSON.stringify(a) !== 1, mes(3))
assert.ok(JSON.stringify(a) == 1, mes(4))

a = 9999999999998888888888887777777777n
assert.equal(JSON.stringify(a), `"9999999999998888888888887777777777"`)

let b = "1"

assert.ok(JSON.stringify(b) === `"1"`, mes(5))
assert.ok(JSON.stringify(b) == `"1"`, mes(6))
assert.ok(JSON.stringify(b) !== 1, mes(7))
assert.ok(JSON.stringify(b) != 1, mes(8))


let c = 1.34
assert.ok(JSON.stringify(c) === "1.34", mes(9))
assert.ok(JSON.stringify(c) == "1.34", mes(10))
assert.ok(JSON.stringify(c) !== 1.34, mes(11))
assert.ok(JSON.stringify(c) == 1.34, mes(12))

let d = true
assert.ok(JSON.stringify(d) === "true", mes(13))
assert.ok(JSON.stringify(d) == "true", mes(14))
assert.ok(JSON.stringify(d) !== true, mes(15))
assert.ok(JSON.stringify(d) == "true", mes(16))

let e = [1, 2, 3]
assert.ok(JSON.stringify(e) === "[1,2,3]", mes(17))
assert.ok(JSON.stringify(e) == "[1,2,3]", mes(18))
assert.ok(JSON.stringify(e) !== [1, 2, 3], mes(19))
assert.ok(JSON.stringify(e) != [1, 2, 3], mes(20))

let f = [1, 2, 3, "a"]
assert.ok(JSON.stringify(f) === `[1,2,3,"a"]`, mes(21))
assert.ok(JSON.stringify(f) == `[1,2,3,"a"]`, mes(22))

let g = {
    a: "1",
    b: 1,
    c: 1.2,
    d: true,
    e: null,
    f: [1, 2, 3]
}
assert.ok(JSON.stringify(g) === `{"a":"1","b":1,"c":1.2,"d":true,"e":null,"f":[1,2,3]}`, mes(23))

assert.ok(JSON.stringify(undefined) === undefined, mes(24))
assert.ok(JSON.stringify(NaN) === "null", mes(25))
assert.ok(JSON.stringify(null) === "null", mes(26))
assert.ok(JSON.stringify(Infinity) === "null", mes(27))