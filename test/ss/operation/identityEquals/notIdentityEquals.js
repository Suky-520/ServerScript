import assert from "assert";

console.log("测试运算符:: !==");

// 数值比较
assert.equal(5 !== 5, false);
assert.equal(5 !== 4, true);
assert.equal(0 !== -0, false); // 注意：在 JavaScript 中，0 和 -0 被认为是相等的

// 数值和字符串
assert.equal(5 !== '5', true); // 不同类型
assert.equal('10' !== '10', false);

// 字符串比较
assert.equal('hello' !== 'hello', false);
assert.equal('hello' !== 'world', true);

// 布尔值比较
assert.equal(true !== true, false);
assert.equal(true !== false, true);

// 特殊数值
assert.equal(NaN !== NaN, true); // 注意：NaN 与任何值都不相等，包括其本身
assert.equal(Infinity !== Infinity, false);
assert.equal(-Infinity !== Infinity, true);

// 不同类型间的比较
assert.equal(0 !== false, true); // 不同类型
assert.equal('1' !== true, true); // 不同类型
assert.equal(null !== undefined, true); // 不同类型

// null 和 undefined
assert.equal(null !== null, false);
assert.equal(undefined !== undefined, false);

// 对象比较
let obj1 = {};
let obj2 = {};
let obj3 = obj1;
assert.equal(obj1 !== obj2, true); // 不同对象实例
assert.equal(obj1 !== obj3, false); // 同一对象实例

// 数组比较
let arr1 = [];
let arr2 = [];
let arr3 = arr1;
assert.equal(arr1 !== arr2, true); // 不同数组实例
assert.equal(arr1 !== arr3, false); // 同一数组实例

// 函数比较
function func1() {}
function func2() {}
assert.equal(func1 !== func2, true); // 不同函数实例
assert.equal(func1 !== func1, false); // 同一函数实例
