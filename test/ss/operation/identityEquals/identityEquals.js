import assert from "assert";

console.log("测试运算符:: ===");

// 数值相等
assert.equal(5 === 5, true);
assert.equal(-1 === -1, true);
assert.equal(0 === 0, true);
assert.equal(-0 === 0, true); // 注意：在 JavaScript 中，+0 和 -0 被认为是严格相等的
assert.equal(0.1 === 0.1, true);

// 数值不等
assert.equal(5 === 10, false);
assert.equal(0 === 1, false);
assert.equal(0.1 === 0.2, false);

// 字符串相等
assert.equal('hello' === 'hello', true);
assert.equal('' === '', true);

// 字符串不等
assert.equal('hello' === 'world', false);
assert.equal('1' === '01', false);

// 布尔值相等
assert.equal(true === true, true);
assert.equal(false === false, true);

// 布尔值不等
assert.equal(true === false, false);

// 特殊数值
assert.equal(NaN === NaN, false); // 注意：NaN 与 NaN 不严格相等
assert.equal(Infinity === Infinity, true);
assert.equal(-Infinity === -Infinity, true);

// 不同类型间的比较
assert.equal(5 === '5', false); // 不同类型
assert.equal('true' === true, false); // 不同类型
assert.equal(0 === false, false); // 不同类型
assert.equal(null === undefined, false); // 不同类型
assert.equal(null === null, true); // 同类型
assert.equal(undefined === undefined, true); // 同类型

// 对象
let obj1 = {};
let obj2 = {};
let obj3 = obj1;
assert.equal(obj1 === obj2, false); // 不同对象实例
assert.equal(obj1 === obj3, true); // 同一对象实例

// 数组
let arr1 = [];
let arr2 = [];
let arr3 = arr1;
assert.equal(arr1 === arr2, false); // 不同数组实例
assert.equal(arr1 === arr3, true); // 同一数组实例

// 函数
function func() {}
assert.equal(func === func, true); // 同一函数实例
