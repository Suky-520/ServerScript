import assert from "assert";

console.log("测试运算符:: <<");

// 基础数值的左移
assert.equal(1 << 2, 4); // 0001 << 2 = 0100，即 4
assert.equal(3 << 3, 24); // 0011 << 3 = 11000，即 24
assert.equal(-1 << 1, -2); // 负数左移，-1 << 1 = -2

// 边界和特殊情况
assert.equal(1 << 0, 1); // 左移 0 位，数值不变
assert.equal(0 << 32, 0); // 左移 32 位，对于 32 位整数，相当于没有移动
assert.equal(1 << 31, -2147483648); // 超过 31 位可能导致符号位改变

// 布尔值的左移
assert.equal(true << 2, 4); // true 转换为 1，然后左移，1 << 2 = 4
assert.equal(false << 3, 0); // false 转换为 0，然后左移，0 << 3 = 0

// 字符串的左移
assert.equal("2" << 2, 8); // 字符串 "2" 转换为数值 2，然后左移，2 << 2 = 8
assert.equal("10" << 1, 20); // 字符串 "10" 转换为数值 10，然后左移，10 << 1 = 20

// 特殊值的左移
assert.equal(null << 2, 0); // null 转换为 0，然后左移，0 << 2 = 0
assert.equal(undefined << 2, 0); // undefined 转换为 NaN，NaN 在数学操作中视为 0，0 << 2 = 0
assert.equal(NaN << 2, 0); // NaN 在数学操作中视为 0，0 << 2 = 0


// 数组的左移（数组转换为原始值）
let arr = [4]; // 数组 toString() 方法返回字符串 "4", 然后转换为数值 4
assert.equal(arr << 2, 16);

// 使用assert.ok来检查布尔值结果
assert.ok((1 << 2) === 4);
assert.ok(("2" << 2) === 8);
