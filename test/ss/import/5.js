import assert from "assert"
/**
 * 测试export 和export def
 */
function mes(i){
    return `language/import/5.js 测试用例[${i}]`
}

import ex,{a as a1,b as b1,c} from './4.js'
assert.deepEqual(ex,{
    name:"44"
},mes(1))
assert.equal(a1,1,mes(2))
assert.equal(b1,2,mes(3))
assert.equal(c,3,mes(4))