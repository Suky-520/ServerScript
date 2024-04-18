import assert from "assert"

console.log("测试 NaN");

assert.equal(isNaN(NaN), true);

// 测试字符串
assert.equal(isNaN("NaN"), true); // 字符串 "NaN"
assert.equal(isNaN("100"), false); // 可转换为数字的字符串
assert.equal(isNaN("Hello"), true); // 不能转换为数字的字符串

// 测试数字
assert.equal(isNaN(100), false); // 普通数字
assert.equal(isNaN(-100), false); // 负数
assert.equal(isNaN(0), false); // 0
assert.equal(isNaN(Infinity), false); // 无穷大
assert.equal(isNaN(-Infinity), false); // 负无穷大

// 测试布尔值
assert.equal(isNaN(true), false); // true 转换为 1
assert.equal(isNaN(false), false); // false 转换为 0

// 测试 undefined 和 null
assert.equal(isNaN(undefined), true); // undefined 转换为 NaN
assert.equal(isNaN(null), false); // null 转换为 0

// 测试对象
assert.equal(isNaN({}), true); // 对象转换为 NaN

// 测试数组
assert.equal(isNaN([]), false); // 空数组转换为 0
assert.equal(isNaN([123]), false); // 单元素数组转换为该元素
assert.equal(isNaN([1, 2]), true); // 多元素数组转换为 NaN
assert.equal(isNaN(["123"]), false); // 单元素数组，元素为字符串，可转换为数字
assert.equal(isNaN(["Hello"]), true); // 单元素数组，元素为字符串，不能转换为数字

// 测试特殊情况
assert.equal(isNaN(new Date()), false); // Date 对象转换为时间戳数字
assert.equal(isNaN(new Date().toString()), true); // Date 转换为字符串，然后无法转换为数字
