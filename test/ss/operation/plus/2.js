import assert from "assert";

console.log("测试运算符:: + [数值 + 字符串 或 字符串 + 数值]");

// 正数 + 字符串
assert.equal(10 + "15", "1015");
assert.equal("20" + 25, "2025");

// 负数 + 字符串
assert.equal(-10 + "15", "-1015");
assert.equal("20" + (-25), "20-25");

// 零 + 字符串
assert.equal(0 + "15", "015");
assert.equal("20" + 0, "200");

// 特殊数值 + 字符串
assert.equal(NaN + "15", "NaN15");
assert.equal("20" + NaN, "20NaN");
assert.equal(Infinity + "15", "Infinity15");
assert.equal("20" + Infinity, "20Infinity");
assert.equal(-Infinity + "15", "-Infinity15");
assert.equal("20" + (-Infinity), "20-Infinity");

// 字符串 + 字符串（模拟数值 + 字符串的特殊情况）
assert.equal("10" + "15", "1015");

// 数值 + 空字符串
assert.equal(10 + "", "10");
assert.equal("" + 20, "20");

// 特殊格式的字符串 + 数值
assert.equal("0x10" + 10, "0x1010"); // 十六进制格式的字符串
assert.equal(10 + "0b101", "100b101"); // 二进制格式的字符串
assert.equal("3.14" + 10, "3.1410"); // 包含小数点的字符串
assert.equal(10 + "3.14", "103.14"); // 数值 + 包含小数点的字符串

// 布尔值 + 字符串（作为附加测试）
assert.equal(true + "10", "true10");
assert.equal("10" + false, "10false");

// null, undefined + 字符串
assert.equal(null + "10", "null10");
assert.equal("10" + null, "10null");
assert.equal(undefined + "10", "undefined10");
assert.equal("10" + undefined, "10undefined");

// 对象、数组 + 字符串
// assert.equal({} + "10", "[object Object]10");
// assert.equal("10" + {}, "10[object Object]");
assert.equal([1,2] + "10", "1,210");
assert.equal("10" + [3,4], "103,4");
