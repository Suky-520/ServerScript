import assert from "assert"

console.log("测试运算符:: |");

// 数值的按位或
assert.equal(1 | 2, 3); // 01 | 10 = 11 (二进制)
assert.equal(5 | 3, 7); // 101 | 011 = 111 (二进制)
assert.equal(-1 | 0, -1); // 负数先转换为二进制补码形式，进行按位或，再转回

// 布尔值与数值的按位或
assert.equal(true | false, 1); // true 转换为 1, false 转换为 0
assert.equal(false | true, 1);

// 字符串与数值的按位或
assert.equal("4" | 1, 5); // 字符串 "4" 转换为数值 4
assert.equal("6" | "3", 7); // 字符串转换为数值进行运算

// 数值与特殊值的按位或
assert.equal(1 | null, 1); // null 转换为 0
assert.equal(1 | undefined, 1); // undefined 转换为 NaN，NaN 视为 0
assert.equal(1 | NaN, 1); // NaN 视为 0

// 数组的按位或（数组转换为原始值）
let arr = [2]; // 数组 toString() 方法返回字符串 "2"，然后转换为数值 2
assert.equal(arr | 3, 3);

// 使用assert.ok来检查布尔值结果
assert.ok((1 | 2) === 3);
