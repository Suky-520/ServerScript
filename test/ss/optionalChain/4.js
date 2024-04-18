import assert from "assert";

console.log("测试语法:: ?.");

let obj = {
    age: 121,
}
let z = obj.name ?? "name";
assert.equal(z, "name");