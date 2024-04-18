import assert from "assert";

console.log("测试运算符:: -");

// 数值减法
assert.equal(100 - 50, 50);
assert.equal(-25 - 25, -50);
assert.equal(0.1 - 0.2, -0.1);

// 数值与字符串
assert.equal("100" - "25", 75);
assert.equal("100" - "a", NaN); // 字符串 "a" 不能转换为数值

// 数值与布尔值
assert.equal(10 - true, 9); // true 被视为 1
assert.equal(10 - false, 10); // false 被视为 0

// 数值与 null
assert.equal(10 - null, 10); // null 被视为 0

// 数值与 undefined
assert.equal(10 - undefined, NaN); // undefined 转换为 NaN

// 字符串减法（字符串可转换为数值时）
assert.equal("50" - "25", 25);
assert.equal("50" - "abc", NaN); // "abc" 不能转换为数值

// 布尔值之间的减法
assert.equal(true - false, 1); // true 转换为 1, false 转换为 0

// 布尔值与 null
assert.equal(false - null, 0); // null 转换为 0

// 布尔值与 undefined
assert.equal(true - undefined, NaN); // undefined 转换为 NaN

// null 与 undefined
assert.equal(null - undefined, NaN); // null 转换为 0, undefined 转换为 NaN

// 特殊数值减法
assert.equal(Infinity - Infinity, NaN);
assert.equal(-Infinity - (-Infinity), NaN);
assert.equal(Infinity - (-Infinity), Infinity);
assert.equal(-Infinity - Infinity, -Infinity);

// 数组减法
// 假设数组转换为其第一个元素的数值（如果仅有一个元素）
assert.equal([100] - [50], 50);
assert.equal([10, 20] - [5, 10], NaN); // 数组转换为逗号分隔的字符串，然后尝试减法

// 使用 assert.strictEqual 和 isNaN 检查 NaN 结果
assert.equal(isNaN("100" - "a"), true);
assert.equal(isNaN(10 - undefined), true);
assert.equal(isNaN("50" - "abc"), true);
assert.equal(isNaN([10, 20] - [5, 10]), true);
