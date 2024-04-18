import assert from "assert";

console.log("测试运算符:: /");

// 数值除法
assert.equal(10 / 5, 2);
assert.equal(-10 / 5, -2);
assert.equal(10 / -5, -2);
assert.equal(10 / 3, 10 / 3); // 无法精确表示的除法
assert.equal(0 / 5, 0);
assert.equal(5 / 0, Infinity); // 任何非零数字除以0为Infinity
assert.equal(-5 / 0, -Infinity); // 负数除以0为-Infinity
assert.equal(0 / 0, NaN); // 0除以0为NaN

// 数值与特殊数值
assert.equal(10 / NaN, NaN);
assert.equal(NaN / 10, NaN);
assert.equal(Infinity / Infinity, NaN);
assert.equal(-Infinity / Infinity, NaN);
assert.equal(Infinity / -Infinity, NaN);
assert.equal(Infinity / 10, Infinity);
assert.equal(10 / Infinity, 0);
assert.equal(-Infinity / 10, -Infinity);
assert.equal(10 / -Infinity, -0);

// 字符串除法（字符串可转换为数值）
assert.equal("10" / "2", 5);
assert.equal("10" / "2.5", 4);
assert.equal("-10" / "2", -5);
assert.equal("10" / "-2", -5);

// 字符串与数值
assert.equal("10" / 2, 5);
assert.equal(10 / "2", 5);
assert.equal("10" / 2.5, 4);
assert.equal(-10 / "2", -5);

// 字符串除法（字符串不可转换为有效数值）
assert.equal("test" / 2, NaN);
assert.equal("10" / "test", NaN);

// 布尔值除法
assert.equal(true / true, 1);
assert.equal(false / true, 0);
assert.equal(true / false, Infinity); // true转换为1，false转换为0

// 布尔值与数值
assert.equal(true / 2, 0.5);
assert.equal(false / 10, 0);

// 布尔值与字符串（可转换为数值）
assert.equal(true / "2", 0.5);
assert.equal(false / "10", 0);

// 布尔值与字符串（不可转换为数值）
assert.equal(true / "test", NaN);
assert.equal(false / "test", NaN);

// null 除法
assert.equal(null / 2, 0);
assert.equal(null / "2", 0);
assert.equal(10 / null, Infinity); // null被视为0

// undefined 除法
assert.equal(undefined / 2, NaN);
assert.equal(undefined / "2", NaN);
assert.equal(10 / undefined, NaN);

// 数组除法
assert.equal([10] / [2], 5);
assert.equal([10, 20] / [2], NaN); // 数组转换为逗号分隔的字符串，然后尝试转换为数值

// 使用assert.ok来检查NaN情况
assert.ok(isNaN("test" / 2));
assert.ok(isNaN("10" / "test"));
assert.ok(isNaN(true / "test"));
assert.ok(isNaN(undefined / 2));
assert.ok(isNaN([10, 20] / [2]));

// 特殊测试用例
assert.equal(1e20 / 1e10, 1e10);

// 循环生成更多测试用例
for (let i = 1; i <= 10; i++) {
    assert.equal(i / i, 1);
    assert.equal(i / -1, -i);
    assert.equal(-i / i, -1);
    assert.equal(i / 0, Infinity);
    assert.equal(-i / 0, -Infinity);
    assert.equal(i / "1", i);
    assert.equal(i / "0", Infinity);
    assert.equal("10" / i, 10 / i);
}
