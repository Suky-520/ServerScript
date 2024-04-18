import assert from "assert"
/**
 * Object测试
 */
function mes(i){
    return `object/object/1.js 测试用例[${i}]`
}

let a={
    name:"xx",
    age:12
}

assert.deepEqual(Object.keys(a),['name','age'],mes(1))
assert.deepEqual(Object.values(a),['xx',12],mes(2))
assert.deepEqual(Object.entries(a),[ [ "name", "xx" ], [ "age", 12 ] ],mes(3))