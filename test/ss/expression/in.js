import assert from "assert";

console.log("测试语法 in");

let a = {
    name: "Xx"
};
let b = [1, 2, 3];

let map = new Map();
map.set("key", 1);
assert.equal("name" in a, true);
assert.equal("name1" in a, false);
assert.equal("0" in b, true);
assert.equal("3" in b, false);
assert.equal("key" in map, false);