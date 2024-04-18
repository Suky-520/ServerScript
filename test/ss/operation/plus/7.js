import assert from "assert";

console.log("测试运算符:: + [空数组 + 空数组]");

assert.equal([] + [], "");

// 单元素数组 + 单元素数组
assert.equal([1] + [2], "12");

// 多元素数组 + 多元素数组
assert.equal([1, 2] + [3, 4], "1,23,4");

// 含有不同类型元素的数组
assert.equal([1, "a"] + [null, true], "1,anull,true");

// 含有复杂对象的数组
// let obj1 = { toString: () => "obj1" };
// let obj2 = { toString: () => "obj2" };
// assert.equal([obj1] + [obj2], "obj1obj2");

// 含有 undefined 和 null 的数组
assert.equal([undefined] + [null], "undefinednull");

// 含有嵌套数组的情况
assert.equal([1, [2, 3]] + [4, [5, 6]], "1,2,34,5,6");

// 特殊数值的情况
assert.equal([NaN] + [Infinity], "NaNInfinity");
assert.equal([-Infinity] + [NaN], "-InfinityNaN");

// 布尔值的情况
assert.equal([true] + [false], "truefalse");

