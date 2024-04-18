import assert from "assert";

console.log("测试运算符:: typeOf");

// 数值
assert.equal(typeOf(42).typeString, 'number');
assert.equal(typeOf(-3.14).typeString, 'number');
assert.equal(typeOf(NaN).typeString, 'number');
assert.equal(typeOf(Infinity).typeString, 'number');

// // 字符串
assert.equal(typeOf('hello').typeString, 'string');
assert.equal(typeOf("").typeString, 'string');

// // 布尔值
assert.equal(typeOf(true).typeString, 'boolean');
assert.equal(typeOf(false).typeString, 'boolean');

// // undefined
assert.equal(typeOf(undefined).typeString, 'undefined');

// // 对象
assert.equal(typeOf({}).typeString, 'object');
assert.equal(typeOf([]).typeString, 'array'); // 注意：JavaScript 中数组被认为是对象类型
assert.equal(typeOf(null).typeString, 'null'); // 注意：这是一个历史遗留问题，null 被认为是 'object'

// 函数
assert.equal(typeOf(function () {
}).typeString, 'function');
assert.equal(typeOf((() => {
})).typeString, 'function');

// BigInt (ES2020 新增)
assert.equal(typeOf(123n).typeString, 'bigint');
assert.equal(typeOf(123n).typeString, 'bigint');

// 使用assert.ok来验证结果是期望的字符串
assert.ok(typeOf(42).typeString === 'number');
assert.ok(typeOf('hello').typeString === 'string');
