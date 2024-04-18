import assert from "assert"

console.log("测试语法:: 变量声明");

function mes(i){
    return `language/variable/1.js 测试用例[${i}]`
}

let a,b =1,f=1
const c=1,d=2,e=3
let g=[1,2,3]
let h= {
    name:"1"
}
assert.equal(a,undefined,mes(1))
assert.equal(b,1,mes(2))
assert.equal(f,1,mes(3))
assert.equal(c,1,mes(4))
assert.equal(d,2,mes(5))
assert.equal(e,3,mes(6))
assert.deepEqual(g,[1,2,3],mes(7))
assert.deepEqual(h,{
    name:"1"
},mes(8))