import assert from "assert";

console.log("测试运算符:: + [undefined + 任何类型]");

// undefined + 数值
assert.equal(undefined + 1, NaN); // undefined 转换为 NaN
assert.equal(undefined + -1, NaN);
assert.equal(undefined + 0, NaN);
assert.equal(undefined + Infinity, NaN);
assert.equal(undefined + -Infinity, NaN);

// undefined + 字符串
assert.equal(undefined + "test", "undefinedtest"); // undefined 转换为 "undefined" 字符串
assert.equal("test" + undefined, "testundefined");

// undefined + 布尔值
assert.equal(undefined + true, NaN); // true 转换为 1, 但 undefined 为 NaN
assert.equal(undefined + false, NaN); // false 转换为 0, 但 undefined 为 NaN

// undefined + null
assert.equal(undefined + null, NaN); // null 转换为 0, 但 undefined 为 NaN

// undefined + undefined
assert.equal(undefined + undefined, NaN); // 两个 undefined 结果为 NaN

// undefined + 对象
// assert.equal(undefined + {}, "undefined[object Object]"); // 对象转换为字符串，undefined 转换为 "undefined"
// assert.equal({} + undefined, "[object Object]undefined");

// undefined + 数组
assert.equal(undefined + [1, 2], "undefined1,2"); // 数组转换为字符串，undefined 转换为 "undefined"
assert.equal([3, 4] + undefined, "3,4undefined");

// 使用assert.ok来检查NaN情况
assert.ok(isNaN(undefined + NaN)); // NaN + undefined
assert.ok(isNaN(undefined + null)); // null + undefined
assert.ok(isNaN(undefined + true)); // true + undefined
assert.ok(isNaN(undefined + false)); // false + undefined
