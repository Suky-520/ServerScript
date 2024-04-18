import assert from "assert";

console.log("测试:: json反序列化");

function mes(i) {
    return `json反序列化 测试用例[${i}]`;
}

let a = "1"
assert.ok(JSON.parse(a) === 1, mes(1))
assert.ok(JSON.parse(a) == 1, mes(2))

let b = "1.3"
assert.ok(JSON.parse(b) === 1.3, mes(3))
assert.ok(JSON.parse(b) == 1.3, mes(4))

let c = "true"
assert.ok(JSON.parse(c) === true, mes(5))
assert.ok(JSON.parse(c) == true, mes(6))

let d = "null"
assert.ok(JSON.parse(d) === null, mes(7))
assert.ok(JSON.parse(d) == null, mes(8))

let e = `"1"`
assert.ok(JSON.parse(e) === "1", mes(9))
assert.ok(JSON.parse(e) == "1", mes(10))
assert.ok(JSON.parse(e) == 1, mes(11))
assert.ok(JSON.parse(e) !== 1, mes(12))

let f = `[1,2,3,"b"]`
assert.deepEqual(JSON.parse(f), [1, 2, 3, "b"], mes(13))

let g = `{"a":"1","b":1,"c":1.2,"d":true,"e":null,"f":[1,2,3]}`
let h = {
    a: "1",
    b: 1,
    c: 1.2,
    d: true,
    e: null,
    f: [1, 2, 3]
};
// assert.deepEqual(JSON.parse(g), h, mes(14));