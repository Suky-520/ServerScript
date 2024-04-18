import assert from "assert"
/**
 * 测试import语法
 */
function mes(i){
    return `language/import/2.js 测试用例[${i}]`
}
import a from "./3.js"

assert.deepEqual(a,{
    name:"33"
},mes(1))

export default {
    name:"2"
}