import assert from "assert";

console.log("测试语法::大整数");

let bu = new Buffer();
assert.ok(typeOf(bu).typeString == "buffer");
bu.write("abc");
assert.ok(bu.size() == 3);