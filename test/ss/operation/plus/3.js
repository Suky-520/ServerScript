import assert from "assert";

console.log("测试运算符:: + [布尔值 + 任何类型]");

// 布尔值 + 数值
assert.equal(true + 1, 2); // true 转换为 1
assert.equal(false + 1, 1); // false 转换为 0
assert.equal(true + -1, 0);
assert.equal(false + -1, -1);

// 布尔值 + 字符串
assert.equal(true + "test", "truetest"); // 布尔值转换为字符串
assert.equal(false + "test", "falsetest");
assert.equal("test" + true, "testtrue");
assert.equal("test" + false, "testfalse");

// 布尔值 + 布尔值
assert.equal(true + true, 2); // 布尔值转换为数值
assert.equal(true + false, 1);
assert.equal(false + false, 0);

// 布尔值 + null
assert.equal(true + null, 1); // null 转换为 0
assert.equal(false + null, 0);

// 布尔值 + undefined
assert.equal(true + undefined+"", "NaN"); // undefined 转换为 NaN
assert.equal(false + undefined+"", "NaN");

// 布尔值 + 特殊数值
assert.equal(true + NaN+"", "NaN");
assert.equal(false + NaN+"", "NaN");
assert.equal(true + Infinity, Infinity);
assert.equal(false + Infinity, Infinity);
assert.equal(true + -Infinity, -Infinity);
assert.equal(false + -Infinity, -Infinity);

// 布尔值 + 对象
// assert.equal(true + {}, "true[object Object]"); // 对象转换为字符串
// assert.equal(false + {}, "false[object Object]");
// assert.equal({} + true, "[object Object]true");
// assert.equal({} + false, "[object Object]false");

// 布尔值 + 数组
assert.equal(true + [1, 2], "true1,2"); // 数组转换为字符串
assert.equal(false + [3, 4], "false3,4");
assert.equal([1, 2] + true, "1,2true");
assert.equal([3, 4] + false, "3,4false");

// 使用assert.ok来检查NaN情况
assert.ok(isNaN(true + undefined));
assert.ok(isNaN(false + undefined));
