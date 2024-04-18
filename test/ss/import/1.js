import assert from "assert"
function mes(i){
    return `language/import/1.js 测试用例[${i}]`
}

import a from "./2.js"

assert.deepEqual(a,{
    name:"2"
},mes(1))