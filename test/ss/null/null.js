import assert from "assert"

console.log("测试语法::空值声明");

// 整型
let a = null;
// 大整数
let b = undefined;

assert.ok(a === null);
assert.ok(b === undefined);