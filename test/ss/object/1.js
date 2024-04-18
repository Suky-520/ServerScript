import assert from "assert"
/**
 * object测试用例
 */

function mes(i){
    return `language/object/1.js 测试用例[${i}]`
}

let a={
    name:"123",
    age:12
}

let b={
    sex:"男",
    name:"bcx",
    "a-1":"a-1",
    ...a,
}

assert.deepEqual(b,{  sex : '男', name : "123",  "a-1" : "a-1",age : 12},mes(1))


let c={
    ...a,
    sex:"男",
    name:"bcx",
    "a-1":"a-1"
}
assert.deepEqual(c,{  name : "bcx",  age : 12, sex : '男',"a-1" : "a-1"},mes(2))
