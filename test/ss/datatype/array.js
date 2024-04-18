import assert from "assert";

console.log("测试基本数据类型::array");

// 正常值测试
let arr = [1, 2, 3, 4, 5];
assert.equal(arr.length, 5, '正常数组的 length 属性不正确');

// 边界值测试
let emptyArr = [];
assert.equal(emptyArr.length, 0, '空数组的 length 属性不为 0');

// join方法
assert.equal(arr.join('-'), '1-2-3-4-5', '正常数组的 join() 方法不正确');
assert.equal(emptyArr.join('-'), '', '空数组的 join() 方法不应返回任何字符');

// push方法
arr.push(6);
assert.equal(arr.length, 6, 'push() 方法未正确增加元素');
assert.equal(arr[arr.length - 1], 6, 'push() 方法未正确增加元素值');
emptyArr.push(1);
assert.equal(emptyArr.length, 1, 'push() 方法未正确在空数组中增加元素');
assert.equal(emptyArr[0], 1, 'push() 方法未正确在空数组中增加元素值');

// 正常值测试
arr = [1, 2, 3, 4, 5];
let poppedElement = arr.pop();
assert.equal(arr.length, 4, 'pop() 方法未正确减少数组长度');
assert.equal(poppedElement, 5, 'pop() 方法未正确返回弹出的元素');

// 边界值测试
emptyArr = [];
let poppedEmptyElement = emptyArr.pop();
assert.equal(emptyArr.length, 0, '空数组 pop() 方法不应影响数组长度');
assert.equal(poppedEmptyElement, undefined, '空数组 pop() 方法应返回 undefined');

// 正常值测试
arr = [1, 2, 3, 4, 5];
let shiftedElement = arr.shift();
assert.equal(arr.length, 4, 'shift() 方法未正确减少数组长度');
assert.equal(shiftedElement, 1, 'shift() 方法未正确返回移除的元素');

// 边界值测试
emptyArr = [];
let shiftedEmptyElement = emptyArr.shift();
assert.equal(emptyArr.length, 0, '空数组 shift() 方法不应影响数组长度');
assert.equal(shiftedEmptyElement, undefined, '空数组 shift() 方法应返回 undefined');

// unshift方法测试
arr = [1, 2, 3, 4, 5];
let lengthBefore = arr.length;
arr.unshift(0);
assert.equal(arr.length, lengthBefore + 1, 'unshift() 方法未正确增加数组长度');
assert.equal(arr[0], 0, 'unshift() 方法未正确在数组开头增加元素');

// 边界值测试
emptyArr = [];
let lengthBeforeEmpty = emptyArr.length;
emptyArr.unshift(1);
assert.equal(emptyArr.length, lengthBeforeEmpty + 1, '空数组 unshift() 方法未正确增加数组长度');
assert.equal(emptyArr[0], 1, '空数组 unshift() 方法未正确在数组开头增加元素');

// 正常值测试
arr = [1, 2, 3, 4, 5];
let sliced = arr.slice(1, 4);
assert.deepEqual(sliced, [2, 3, 4], 'slice() 方法未正确返回期望的子数组');

// 边界值测试
emptyArr = [];
let slicedEmpty = emptyArr.slice(1, 4);
assert.deepEqual(slicedEmpty, [], '空数组 slice() 方法未正确返回空数组');

// 正常值测试
arr = [1, 2, 3, 4, 5];
let spliced = arr.splice(1, 2);
assert.deepEqual(arr, [1, 4, 5], 'splice() 方法未正确删除数组元素');
assert.deepEqual(spliced, [2, 3], 'splice() 方法未正确返回期望的被删除的元素');
arr = [1, 2, 3, 4, 5];
spliced = arr.splice(1,);
assert.deepEqual(arr, [1], 'splice() 方法未正确删除数组元素');
assert.deepEqual(spliced, [2, 3, 4, 5], 'splice() 方法未正确返回期望的被删除的元素');

// 边界值测试
emptyArr = [];
let splicedEmpty = emptyArr.splice(1, 2);
assert.deepEqual(emptyArr, [], '空数组 splice() 方法未正确处理空数组');
assert.deepEqual(splicedEmpty, [], '空数组 splice() 方法未正确返回空数组');

// 正常值测试
arr = [1, 2, 3, 4, 5];
assert.equal(arr.indexOf(3), 2, 'indexOf() 方法未正确返回元素索引');
assert.equal(arr.indexOf(6), -1, 'indexOf() 方法未正确处理不存在的元素');

// 边界值测试
emptyArr = [];
assert.equal(emptyArr.indexOf(1), -1, '空数组 indexOf() 方法未正确处理空数组');

// 正常值测试
arr = [1, 2, 3, 4, 3, 5];
assert.equal(arr.lastIndexOf(3), 4, 'lastIndexOf() 方法未正确返回最后一个匹配元素的索引');
assert.equal(arr.lastIndexOf(6), -1, 'lastIndexOf() 方法未正确处理不存在的元素');

// 边界值测试
emptyArr = [];
assert.equal(emptyArr.lastIndexOf(1), -1, '空数组 lastIndexOf() 方法未正确处理空数组');

// 正常值测试
arr = [1, 2, 3, 4, 5];
let filtered = arr.filter(num => num % 2 === 0);
assert.deepEqual(filtered, [2, 4], 'filter() 方法未正确过滤数组元素');

// 边界值测试
emptyArr = [];
let filteredEmpty = emptyArr.filter(num => num > 0);
assert.deepEqual(filteredEmpty, [], '空数组 filter() 方法未正确处理空数组');

// 字符串排序
let nums = [72, 28, 100, 46, 15, 10, 27];
let arrStr = ["banana", "apple", "pear", "orange"];
nums.sort();
assert.deepEqual(nums, [10, 100, 15, 27, 28, 46, 72], '数组 sort() 方法未正确排序');
arrStr.sort();
assert.deepEqual(arrStr, ["apple", "banana", "orange", "pear"], '数组 sort() 方法未正确排序');

// 数字排序
let numList = [72, 28, 100, 46, 15, 10, 27];
numList.sort((a, b) => {
    return a - b;
});
assert.deepEqual(numList, [10, 15, 27, 28, 46, 72, 100], '数组 sort() 方法未正确排序');
numList.sort((a, b) => {
    return b - a;
});
assert.deepEqual(numList, [100, 72, 46, 28, 27, 15, 10], '数组 sort() 方法未正确排序');

// 复杂排序
let userList = [
    {age: 20, name: "ZhangSan"},
    {age: 19, name: "AoDuo"},
    {age: 18, name: "LiSi"},
    {age: 22, name: "WangWu"},
    {age: 20, name: "ZhaoLiu"}
]
userList.sort((u1, u2) => {
    return u1.age - u2.age;
});
assert.deepEqual(userList, [
    {age: 18, name: "LiSi"},
    {age: 19, name: "AoDuo"},
    {age: 20, name: "ZhangSan"},
    {age: 20, name: "ZhaoLiu"},
    {age: 22, name: "WangWu"}]);
userList.sort((u1, u2) => {
    return u1.name.localeCompare(u2.name);
});
assert.deepEqual(userList, [
    {age: 19, name: "AoDuo"},
    {age: 18, name: "LiSi"},
    {age: 22, name: "WangWu"},
    {age: 20, name: "ZhangSan"},
    {age: 20, name: "ZhaoLiu"}]);
