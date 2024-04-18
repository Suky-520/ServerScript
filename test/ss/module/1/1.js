import assert from "assert"
import tt from './2.js'
assert.ok(tt().getName().getAge()=="vvv")
assert.ok(tt().getAge()=="vino")