import assert from "assert"

console.log("测试运算符:: >>>")

// 基础数值的无符号右移
assert.equal(4 >>> 2, 1); // 0100 >>> 2 = 0001，即 1
assert.equal(-4 >>> 2, 1073741823); // 负数无符号右移，进行补零
assert.equal(256 >>> 4, 16); // 100000000 >>> 4 = 00010000，即 16

// 边界和特殊情况
assert.equal(1 >>> 0, 1); // 无符号右移 0 位，数值不变
assert.equal(-1 >>> 31, 1); // -1 无符号右移 31 位，最高位变为 0，其余位为 1

// 布尔值的无符号右移
assert.equal(true >>> 1, 0); // true 转换为 1，然后无符号右移，1 >>> 1 = 0
assert.equal(false >>> 1, 0); // false 转换为 0，然后无符号右移，0 >>> 1 = 0

// 字符串的无符号右移
assert.equal("8" >>> 2, 2); // 字符串 "8" 转换为数值 8，然后无符号右移，8 >>> 2 = 2
assert.equal("1024" >>> 3, 128); // 字符串 "1024" 转换为数值 1024，然后无符号右移，1024 >>> 3 = 128

// 特殊值的无符号右移
assert.equal(null >>> 2, 0); // null 转换为 0，然后无符号右移，0 >>> 2 = 0
assert.equal(undefined >>> 2, 0); // undefined 转换为 NaN，NaN 在数学操作中视为 0，0 >>> 2 = 0
assert.equal(NaN >>> 2, 0); // NaN 在数学操作中视为 0，0 >>> 2 = 0

// 数组的无符号右移（数组转换为原始值）
let arr = [128]; // 数组 toString() 方法返回字符串 "128", 然后转换为数值 128
assert.equal(arr >>> 3, 16);

// 使用assert.ok来检查布尔值结果
assert.ok((4 >>> 2) === 1);
assert.ok(("8" >>> 2) === 2);
