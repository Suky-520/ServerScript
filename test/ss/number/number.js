import assert from "assert"

console.log("测试语法::数字声明");

// 整型
let a = 12;
// 损失精度
let b = 9223372036854775808; //结果：9.223372036854776e+18
// 浮点
let c = 1.23;

assert.ok(a === 12);
assert.ok(b === 9223372036854775808);
assert.ok(c === 1.23);