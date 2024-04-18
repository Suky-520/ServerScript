import assert from "assert"

console.log("测试运算符:: &&");

// 布尔值逻辑与
assert.equal(true && true, true);
assert.equal(true && false, false);
assert.equal(false && true, false);
assert.equal(false && false, false);

// 数值逻辑与
assert.equal(1 && 2, 2); // 两个真值，返回第二个操作数
assert.equal(0 && 2, 0); // 第一个假值，返回第一个操作数
assert.equal(1 && 0, 0); // 第二个假值，返回第二个操作数
assert.equal(-1 && 2, 2); // 负数被视为真值

// 字符串逻辑与
assert.equal('hello' && 'world', 'world'); // 两个非空字符串
assert.equal('' && 'world', ''); // 空字符串被视为假值
assert.equal('hello' && '', ''); // 空字符串被视为假值

// null 和 undefined 逻辑与
assert.equal(null && 'hello', null);
assert.equal('hello' && null, null);
assert.equal(undefined && 'world', undefined);
assert.equal('world' && undefined, undefined);

// 数组和对象逻辑与
assert.equal([] && 'hello', 'hello'); // 空数组被视为真值
assert.equal({} && 'world', 'world'); // 空对象被视为真值

// 函数逻辑与
function myFunc() {
    return 'hello';
}

assert.equal(myFunc() && 'world', 'world'); // 函数调用返回值与字符串

// 混合类型逻辑与
assert.equal(1 && 'hello', 'hello'); // 数值和字符串
assert.equal(false && {}, false); // 布尔值和对象
assert.deepEqual('yes' && [1, 2, 3], [1, 2, 3]); // 字符串和数组

// 特殊数值逻辑与
assert.equal(NaN && 'hello', NaN); // NaN 被视为假值
assert.equal('hello' && NaN, NaN); // NaN 被视为假值
assert.equal(Infinity && 'world', 'world'); // Infinity 被视为真值
assert.equal('world' && Infinity, Infinity); // Infinity 被视为真值

// 使用assert.ok来检查布尔值结果
assert.ok(true && true);
assert.equal((false && true), false);
