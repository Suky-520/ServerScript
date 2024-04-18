import assert from "assert";

console.log("测试运算符:: + [数值 + 数值]");

// 正数 + 正数
assert.equal(10 + 15, 25);

// 正数 + 负数
assert.equal(10 + (-5), 5);

// 负数 + 负数
assert.equal((-10) + (-10), -20);

// 正数 + 0
assert.equal(10 + 0, 10);

// 负数 + 0
assert.equal((-10) + 0, -10);

// 0 + 0
assert.equal(0 + 0, 0);

// 正无穷大 + 正无穷大
assert.equal(Infinity + Infinity, Infinity);

// 负无穷大 + 负无穷大
assert.equal(-Infinity + -Infinity, -Infinity);

// 正无穷大 + 负无穷大
assert.ok(isNaN(Infinity + -Infinity));

// 正无穷大 + 有限数
assert.equal(Infinity + 10, Infinity);

// 负无穷大 + 有限数
assert.equal(-Infinity + 10, -Infinity);

// 正无穷大 + 0
assert.equal(Infinity + 0, Infinity);

// 负无穷大 + 0
assert.equal(-Infinity + 0, -Infinity);

// NaN + 任何数值
assert.ok(isNaN(NaN + 10));

// 任何数值 + NaN
assert.ok(isNaN(10 + NaN));

// NaN + NaN
assert.ok(isNaN(NaN + NaN));

// 使用浮点数
assert.equal(0.1 + 0.2, 0.30000000000000004);
assert.equal(1.5 + 2.5, 4.0);

// 特大或特小的数值
assert.equal(1e308 + 1e308, Infinity); // 超出最大数值范围
assert.equal(-1e308 + -1e308, -Infinity); // 超出最小数值范围
