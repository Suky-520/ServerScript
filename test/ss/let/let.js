import assert from "assert"

console.log("测试语法::变量声明");

// 变量声明
let str = "hello, world!";
let num = 1;
let bigint = 123456789012345678901234567890n;
let bool = true;
let nil = null;
let undefinedObj = undefined;

assert.ok(str === "hello, world!");
assert.ok(num === 1);
assert.ok(bigint === 123456789012345678901234567890n);
assert.ok(bool === true);
assert.ok(nil === null);
assert.ok(undefinedObj === undefined);


// 常量声明
const str2 = "hello, world!";
const num2 = 1;
const bigint2 = 123456789012345678901234567890n;
const bool2 = true;
const nil2 = null;
const undefinedObj2 = undefined;

assert.ok(str2 === "hello, world!");
assert.ok(num2 === 1);
assert.ok(bigint2 === 123456789012345678901234567890n);
assert.ok(bool2 === true);
assert.ok(nil2 === null);
assert.ok(undefinedObj2 === undefined);