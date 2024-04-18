import assert from "assert";

console.log("测试运算符:: ==");

// 数值和字符串
assert.ok(0 == '0');
assert.ok(1 == '1');
assert.ok(-1 == '-1');
assert.ok(0.1 == '0.1');

// 布尔值和其他类型
assert.ok(true == 1);
assert.ok(false == 0);
assert.ok(true == '1');
assert.ok(false == '0');

let arr = [1];
assert.ok(arr == '1'); // 数组转换为字符串

// null 和 undefined
assert.ok(null == undefined);
assert.equal(null == 0,false);
assert.equal(undefined == 0,false);

// 数值比较
assert.equal(NaN == NaN,false); // NaN 与任何值都不相等，包括其本身
assert.ok(0 == 0);
assert.ok(-0 == 0); // -0 和 +0 是相等的
assert.ok(Infinity == Infinity);
assert.ok(-Infinity == -Infinity);

// 字符串比较
assert.ok('hello' == 'hello');
assert.equal('hello' == 'world',false);

// 混合类型比较
assert.equal(0 == '',true);
assert.equal(0 == ' ',false);
assert.ok(0 == false);
assert.ok('1' == true);
assert.ok('0' == false);
assert.equal(1 == true,true);
assert.equal('true' == true,false); // false，因为布尔值 true 转换为 1
assert.equal('false' == false,false); // false，因为布尔值 false 转换为 0

// 特殊数值
assert.equal(Infinity == -Infinity,false);
assert.equal(Infinity == NaN,false);
assert.equal(NaN == 0,false);
assert.equal(NaN == 'NaN',false);

assert.ok([3] == 3);
assert.ok([3] == '3');
assert.equal([2, 3] == '2,3',true); // true，因为数组转换为字符串 "2,3"
//assert.equal({} == '[object Object]',true); // false，尽管对象转换为字符串 '[object Object]'，但隐式转换不适用于此类直接比较

// 生成更多测试用例
for (let i = 0; i <= 10; i++) {
    assert.ok(i.toString() == i);
    assert.equal(((i - 10) == (i + '0') / 10 - 1),false);
    assert.ok((i % 2 == 0) == !(i % 2));
    assert.ok((i % 2 == 1) == !!(i % 2));
}


// 对象比较
let obj1 = {};
let obj2 = {};
let obj3 = obj1;
assert.equal(obj1 == obj2, false); // 不同对象实例
assert.equal(obj1 == obj3, true); // 同一对象实例


// 数组比较
let arr1 = [];
let arr2 = [];
let arr3 = arr1;
assert.equal(arr1 == arr2, false); // 不同数组实例
assert.equal(arr1 == arr3, true); // 同一数组实例

// 函数比较
function func1() {}
function func2() {}
assert.equal(func1 == func2, false); // 不同函数实例
assert.equal(func1 == func1, true); // 同一函数实例
