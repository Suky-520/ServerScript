import assert from "assert";

console.log("测试运算符:: + [null + 任何类型]");

// null + 数值
assert.equal(null + 1, 1); // null 转换为 0
assert.equal(null + -1, -1);
assert.equal(null + 0, 0);
assert.equal(null + Infinity, Infinity);
assert.equal(null + -Infinity, -Infinity);
assert.equal(null + NaN+"", "NaN"); // NaN 情况

// null + 字符串
assert.equal(null + "test", "nulltest"); // null 转换为 "null" 字符串
assert.equal("test" + null, "testnull");

// null + 布尔值
assert.equal(null + true, 1); // true 转换为 1
assert.equal(null + false, 0); // false 转换为 0

// null + null
assert.equal(null + null, 0); // 两个 null 都转换为 0

// null + undefined
assert.equal(null + undefined+"", "NaN"); // 结果为 NaN

// null + 对象
// assert.equal(null + {}, "null[object Object]"); // 对象转换为字符串
// assert.equal({} + null, "[object Object]null");

// null + 数组
assert.equal(null + [1, 2], "null1,2"); // 数组转换为字符串
assert.equal([3, 4] + null, "3,4null");

// 使用assert.ok来检查NaN情况
assert.ok(isNaN(null + NaN));
assert.ok(isNaN(null + undefined));

