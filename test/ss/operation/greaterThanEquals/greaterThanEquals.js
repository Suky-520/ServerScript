import assert from "assert";

console.log("测试运算符:: >=");

// 数值比较
assert.equal(5 >= 3, true);
assert.equal(3 >= 3, true);
assert.equal(-1 >= -2, true);
assert.equal(0.5 >= 0.1, true);
assert.equal(2 >= 3, false);

// 字符串比较（按字典顺序）
assert.equal('b' >= 'a', true);
assert.equal('abc' >= 'ab', true);
assert.equal('12' >= '2', false); // 字符串比较，'1' 在字典顺序上小于 '2'

// 数值与字符串
assert.equal(10 >= '2', true); // 字符串 '2' 转换为数值 2
assert.equal('5' >= 10, false); // 字符串 '5' 转换为数值 5

// 布尔值
assert.equal(true >= false, true); // true 转换为 1, false 转换为 0
assert.equal(false >= true, false);

// null 和 undefined
assert.equal(null >= 0, true); // null 转换为 0
assert.equal(null >= 1, false); // null 转换为 0
assert.equal(undefined >= 0, false); // undefined 转换为 NaN

// 数组和对象（通过转换为原始值比较）
assert.equal([2] >= 1, true); // 数组 [2] 转换为字符串 "2"，然后转换为数值 2
assert.equal([3] >= '2', true); // 数组 [3] 转换为字符串 "3"
assert.equal({} >= 0, false); // 对象 {} 无法转换为一个有意义的数值进行比较

// 特殊数值
assert.equal(NaN >= NaN, false);
assert.equal(5 >= Infinity, false);
assert.equal(-Infinity >= -5, false);
assert.equal(Infinity >= 5, true);

// 使用assert.equal来检查布尔结果
assert.equal(3 >= 5, false);
assert.equal('b' >= 'a', true);
assert.equal(true >= false, true);
assert.equal(null >= 0, true);
assert.equal(undefined >= 0, false);
assert.equal([2] >= 1, true);
assert.equal(NaN >= NaN, false);
assert.equal(Infinity >= 5, true);

// 注意：NaN、undefined 和对象比较的行为
assert.equal(NaN >= 3, false);
assert.equal(undefined >= '5', false);
assert.equal({} >= 0, false); // 对象与数值比较通常返回 false 除非对象定义了 valueOf 或 toString 方法
