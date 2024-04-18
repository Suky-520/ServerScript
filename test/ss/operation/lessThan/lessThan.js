import assert from "assert";

console.log("测试运算符:: <");

// 数值比较
assert.equal(3 < 5, true);
assert.equal(3 < 3, false);
assert.equal(-2 < -1, true);
assert.equal(0.1 < 0.5, true);

// 字符串比较（按字典顺序）
assert.equal('a' < 'b', true);
assert.equal('ab' < 'abc', true);
assert.equal('12' < '2', true); // 字符串比较，'1' 在字典顺序上小于 '2'

// 数值与字符串
assert.equal(2 < '10', true); // 字符串 '10' 转换为数值 10
assert.equal('5' < 10, true); // 字符串 '5' 转换为数值 5

// 布尔值
assert.equal(false < true, true); // false 转换为 0, true 转换为 1
assert.equal(true < false, false);

// null 和 undefined
assert.equal(null < 0, false); // null 转换为 0
assert.equal(null < 1, true); // null 转换为 0
assert.equal(undefined < 0, false); // undefined 转换为 NaN

// 数组和对象（通过转换为原始值比较）
assert.equal([1] < 2, true); // 数组 [1] 转换为字符串 "1"，然后转换为数值 1
assert.equal([3] < '2', false); // 数组 [3] 转换为字符串 "3"
let ob={}
assert.equal(ob<0, false); // 对象 {} 无法转换为一个有意义的数值进行比较

// 特殊数值
assert.equal(NaN < NaN, false);
assert.equal(5 < Infinity, true);
assert.equal(-Infinity < -5, true);

// 使用assert.equal来检查布尔结果
assert.equal(3 < 5, true);
assert.equal('a' < 'b', true);
assert.equal(false < true, true);
assert.equal(null < 0, false);
assert.equal(undefined < 0, false);
assert.equal([1] < 2, true);
assert.equal(NaN < NaN, false);
assert.equal(5 < Infinity, true);

// 注意：NaN、undefined 和对象比较的行为
assert.equal(NaN < 3, false);
assert.equal(undefined < '5', false);
assert.equal(ob < 0, false); // 对象与数值比较通常返回 false 除非对象定义了 valueOf 或 toString 方法
