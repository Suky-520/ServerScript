import assert from "assert";

console.log("测试运算符:: *");

// 数值乘法
assert.equal(2 * 3, 6);
assert.equal(-2 * 3, -6);
assert.equal(2.5 * 4, 10);
assert.equal(-2.5 * -4, 10);
assert.equal(0 * 100, 0);

// 数值与特殊数值
assert.equal(2 * NaN, NaN);
assert.equal(Infinity * 0, NaN);
assert.equal(-Infinity * 0, NaN);
assert.equal(Infinity * Infinity, Infinity);
assert.equal(-Infinity * Infinity, -Infinity);
assert.equal(Infinity * -Infinity, -Infinity);

// 字符串乘法（字符串可转换为数值）
assert.equal("2" * "3", 6);
assert.equal("2.5" * "4", 10);
assert.equal("-2" * "-3", 6);

// 字符串与数值
assert.equal("2" * 3, 6);
assert.equal(2 * "3", 6);
assert.equal("2.5" * 4, 10);
assert.equal(-2 * "-3", 6);

// 字符串乘法（字符串不可转换为有效数值）
assert.equal("test" * 2, NaN);
assert.equal("two" * "three", NaN);

// 布尔值乘法
assert.equal(true * true, 1);
assert.equal(false * true, 0);
assert.equal(true * false, 0);
assert.equal(false * false, 0);

// 布尔值与数值
assert.equal(true * 2, 2);
assert.equal(false * 10, 0);

// 布尔值与字符串（可转换为数值）
assert.equal(true * "2", 2);
assert.equal(false * "10", 0);

// 布尔值与字符串（不可转换为数值）
assert.equal(true * "test", NaN);
assert.equal(false * "test", NaN);

// null 乘法
assert.equal(null * 2, 0);
assert.equal(null * "2", 0);
assert.equal(null * null, 0);

// undefined 乘法
assert.equal(undefined * 2, NaN);
assert.equal(undefined * "2", NaN);
assert.equal(undefined * undefined, NaN);


// 数组乘法
assert.equal([2] * 3, 6);
assert.equal([2, 3] * [2], NaN); // 数组转换为字符串，然后尝试转换为数值
assert.equal(["2"] * 3, 6);
assert.equal(["two"] * 3, NaN);

// 使用assert.ok来检查NaN情况，因为NaN !== NaN
assert.ok(isNaN("test" * 2));
assert.ok(isNaN("two" * "three"));
assert.ok(isNaN(true * "test"));
assert.ok(isNaN(undefined * 2));
assert.ok(isNaN([2, 3] * [2]));
assert.ok(isNaN(["two"] * 3));

// 特殊测试用例
assert.equal(0.1 * 0.2, 0.020000000000000004); // 浮点数精度问题
assert.equal(1e20 * 1e20, 1e40); // 大数乘法

// 循环生成更多测试用例
for (let i = 1; i <= 10; i++) {
    assert.equal(i * i, Math.pow(i, 2));
    assert.equal(i * -1, -i);
    assert.equal(i * 0, 0);
    assert.equal(i * 1, i);
}
