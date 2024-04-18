import assert from "assert"
let a={
    name:"X",
    age:10,
}
let c = delete a.age
assert.ok(c===true)
assert.deepEqual(a,{  name : "X"})
