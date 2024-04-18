import assert from "assert"

console.log("测试运算符:: +");

// 正号操作符测试用例
assert.equal(+42, 42); // 数值
assert.equal(+'42', 42); // 字符串转换为数值
assert.equal(+true, 1); // 布尔值 true 转换为 1
assert.equal(+false, 0); // 布尔值 false 转换为 0
assert.equal(+'', 0); // 空字符串转换为 0
assert.equal(+'hello', NaN); // 非数值字符串转换为 NaN
assert.equal(+null, 0); // null 转换为 0
assert.equal(+undefined, NaN); // undefined 转换为 NaN

// 负号操作符测试用例
assert.equal(-42, -42); // 数值
assert.equal(-'42', -42); // 字符串转换为数值并取反
assert.equal(-true, -1); // 布尔值 true 转换为 1 并取反
assert.equal(-false, -0); // 布尔值 false 转换为 0
assert.equal(-'', -0); // 空字符串转换为 0
assert.equal(-'hello', NaN); // 非数值字符串转换为 NaN
assert.equal(-null, -0); // null 转换为 0
assert.equal(-undefined, NaN); // undefined 转换为 NaN

// 使用assert.ok来验证结果是否符合预期（对于 NaN 的检查）
assert.ok(isNaN(+'hello'));
assert.ok(isNaN(-'hello'));
