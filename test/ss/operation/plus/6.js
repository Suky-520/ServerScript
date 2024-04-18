import assert from "assert";

console.log("测试运算符:: + [对象（包括数组）+ 任何类型]");

// 对象 + 数值
// let obj = { toString: () => "3" };
// assert.equal(obj + 2, "32"); // 对象转换为字符串 "3"，然后与数字 2 进行字符串连接
//
// // 对象 + 字符串
// obj = { toString: () => "object" };
// assert.equal(obj + " test", "object test"); // 对象转换为字符串 "object"，然后与 " test" 字符串连接
//
// // 对象 + 布尔值
// obj = { toString: () => "true", valueOf: () => 1 };
// assert.equal(obj + true, 2); // 对象通过 valueOf 转换为数字 1，然后与 true（转换为 1）相加
//
// // 对象 + null
// obj = { toString: () => "nullObject" };
// assert.equal(obj + null, "nullObjectnull"); // 对象转换为字符串 "nullObject"，然后与 "null" 字符串连接
//
// // 对象 + undefined
// obj = { toString: () => "undefinedObject" };
// assert.equal(obj + undefined, "undefinedObjectundefined"); // 对象转换为字符串 "undefinedObject"，然后与 "undefined" 字符串连接
//
// // 对象 + 对象
// let obj1 = { toString: () => "obj1" };
// let obj2 = { toString: () => "obj2" };
// assert.equal(obj1 + obj2, "obj1obj2"); // 两个对象分别转换为它们的字符串表示，然后进行连接

// 数组 + 数值
assert.equal([1, 2] + 3, "1,23"); // 数组转换为字符串 "1,2"，然后与数字 3 进行字符串连接

// 数组 + 字符串
assert.equal([1, 2] + " test", "1,2 test"); // 数组转换为字符串 "1,2"，然后与 " test" 字符串连接

// 数组 + 布尔值
assert.equal([1] + true, "1true"); // 数组转换为字符串 "1"，然后与 "true" 字符串连接

// 数组 + null
assert.equal([1, 2] + null, "1,2null"); // 数组转换为字符串 "1,2"，然后与 "null" 字符串连接

// 数组 + undefined
assert.equal([1, 2] + undefined, "1,2undefined"); // 数组转换为字符串 "1,2"，然后与 "undefined" 字符串连接

// 数组 + 数组
assert.equal([1, 2] + [3, 4], "1,23,4"); // 两个数组分别转换为它们的字符串表示，然后进行连接

