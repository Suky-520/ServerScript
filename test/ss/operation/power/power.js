import assert from "assert";

console.log("测试运算符:: **");

// 正数指数
assert.equal(2 ** 3, 8);
assert.equal(3 ** 2, 9);
assert.equal(2 ** 0, 1);
assert.equal(2 ** -1, 0.5);
assert.equal(-(2 ** 2), -4); // 注意：由于优先级，实际为-(2**2)
assert.equal((-2) ** 2, 4);
assert.equal((-2) ** 3, -8);

// 零和负数指数
assert.equal(0 ** 1, 0);
assert.equal(0 ** 0, 1); // 数学上未定义，但在JS中为1
assert.equal(1 ** 0, 1);
assert.equal((-1) ** 0, 1);
assert.equal(0 ** -1, Infinity);

// 特殊数值
assert.equal(Infinity ** 0, 1);
assert.equal(Infinity ** 2, Infinity);
assert.equal(Infinity ** -1, 0);
assert.equal((-Infinity) ** 2, Infinity);
assert.equal((-Infinity) ** 3, -Infinity);
assert.equal(NaN ** 2, NaN);
assert.equal(2 ** NaN, NaN);

// 与字符串的操作
assert.equal(2 ** "3", 8);
assert.equal("2" ** 3, 8);
assert.equal("2" ** "3", 8);
assert.equal(2 ** "-1", 0.5);
assert.equal("-2" ** 2, 4); // 注意：由于优先级，实际为-(2**2)
assert.equal("2" ** "a", NaN); // "a" 无法转换为数值

// 与布尔值的操作
assert.equal(2 ** true, 2);
assert.equal(2 ** false, 1);
assert.equal(true ** 2, 1); // true 转换为 1
assert.equal(false ** 2, 0); // false 转换为 0

// 与 null 的操作
assert.equal(2 ** null, 1); // null 转换为 0
assert.equal(null ** 2, 0); // null 转换为 0

// 与 undefined 的操作
assert.equal(2 ** undefined, NaN);
assert.equal(undefined ** 2, NaN);

// 数组求指数
assert.equal([2] ** [3], 8);
assert.equal([2, 3] ** [2], NaN); // 数组转换为逗号分隔的字符串，然后尝试转换为数值

// 使用assert.ok来检查NaN情况
assert.ok(isNaN("2" ** "a"));
assert.ok(isNaN(2 ** undefined));
assert.ok(isNaN([2, 3] ** [2]));

// 循环生成更多测试用例
for (let i = 1; i <= 10; i++) {
    assert.equal(i ** 2, Math.pow(i, 2));
    assert.equal(i ** 0, 1);
    assert.equal(i ** 1, i);
    assert.equal((-i) ** 2, Math.pow(-i, 2));
    assert.equal(2 ** (i / 10), 2 ** (i / 10)); // 比较浮点数指数
}
