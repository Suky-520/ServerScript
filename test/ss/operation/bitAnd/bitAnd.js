import assert from "assert"

console.log("测试运算符:: &");

// 数值的按位与
assert.equal(5 & 3, 1); // 0101 & 0011 = 0001
assert.equal(-2 & -3, -4); // 负数先转换为二进制补码形式，进行按位与，再转回
assert.equal(0xFFFF & 0x0000, 0); // 全1与全0 = 0

// 布尔值与数值的按位与
assert.equal(true & 1, 1); // true 转换为 1
assert.equal(false & 1, 0); // false 转换为 0

// 字符串与数值的按位与
assert.equal("4" & 5, 4); // 字符串 "4" 转换为数值 4
assert.equal("6" & "3", 2); // 字符串转换为数值进行运算

// 数值与特殊值的按位与
assert.equal(1 & null, 0); // null 转换为 0
assert.equal(1 & undefined, 0); // undefined 转换为 NaN，NaN 视为 0
assert.equal(1 & NaN, 0); // NaN 视为 0

// 数组的按位与（数组转换为原始值）
let arr = [2]; // 数组 toString() 方法返回字符串 "2"，然后转换为数值 2
assert.equal(arr & 3, 2);

// 进行更复杂的按位与运算
assert.equal((15 & 9) & 3, 1); // 先 15 & 9 得到 9，再与 3 进行按位与得到 1

// 使用assert.ok来检查布尔值结果
assert.ok((5 & 3) === 1);
