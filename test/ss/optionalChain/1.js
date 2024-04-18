import assert from "assert";

console.log("测试语法:: ?.");

const obj4 = null;
assert.equal(obj4?.method(), undefined)

// 测试访问对象属性
const obj1 = {prop: 42};
assert.equal(obj1?.prop, 42);

const obj2 = null;
assert.equal(obj2?.prop, undefined);

// 测试访问对象方法
const obj3 = {method: () => 42};
assert.equal(obj3?.method(), 42);

// 测试访问数组元素
const arr1 = [1, 2, 3];
assert.equal(arr1?.[0], 1);
const arr2 = null;
assert.equal(arr2?.[0], undefined);

// 测试多级可选链
const nestedObj1 = {outer: {inner: 42}};
assert.equal(nestedObj1?.outer?.inner, 42);
const nestedObj2 = null;
assert.equal(nestedObj2?.outer?.inner, undefined);

// 测试混合使用可选链和普通链
const mixedObj = {prop: {method: () => 42}};
assert.equal(mixedObj?.prop.method(), 42);
const mixedObj2 = null;
assert.equal(mixedObj2?.prop.method(), undefined);


