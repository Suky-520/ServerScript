import assert from "assert";

console.log("测试运算符:: ~");

// 数值的按位非
assert.equal(~0, -1); // ~0000 = 1111 (二进制), 二进制补码表示 -1
assert.equal(~1, -2); // ~0001 = 1110 (二进制), 二进制补码表示 -2
assert.equal(~-1, 0); // ~(-1) = 0, 因为 -1 的二进制补码表示全是 1
assert.equal(~5, -6); // ~0101 = 1010 (二进制), 二进制补码表示 -6

// 布尔值的按位非
assert.equal(~true, -2); // true 转换为 1, ~1 = -2
assert.equal(~false, -1); // false 转换为 0, ~0 = -1

// 字符串的按位非
assert.equal(~"2", -3); // 字符串 "2" 转换为数值 2, ~2 = -3
assert.equal(~"0", -1); // 字符串 "0" 转换为数值 0, ~0 = -1

// 特殊值的按位非
assert.equal(~null, -1); // null 转换为 0, ~0 = -1
assert.equal(~undefined, -1); // undefined 转换为 NaN, NaN 在按位操作中视为 0, ~0 = -1
assert.equal(~NaN, -1); // NaN 在按位操作中视为 0, ~0 = -1

// 数组的按位非（数组转换为原始值）
let arr = [2]; // 数组 toString() 方法返回字符串 "2", 然后转换为数值 2
assert.equal(~arr, -3);

// 使用assert.ok来检查布尔值结果
assert.ok(~0 === -1);
assert.ok(~-1 === 0);