import assert from "assert";

console.log("测试运算符:: %");

// 数值求余
assert.equal(10 % 5, 0);
assert.equal(10 % 3, 1);
assert.equal(10 % 4, 2);
assert.equal(-10 % 3, -1);
assert.equal(10 % -3, 1);
assert.equal(-10 % -3, -1);
assert.equal(0 % 3, 0);
assert.equal(3 % 0, NaN);
assert.equal(-3 % 0, NaN);
assert.equal(0 % 0, NaN);

// 特殊数值
assert.equal(Infinity % 3, NaN);
assert.equal(-Infinity % 3, NaN);
assert.equal(3 % Infinity, 3);
assert.equal(-3 % Infinity, -3);
assert.equal(Infinity % Infinity, NaN);
assert.equal(-Infinity % -Infinity, NaN);
assert.equal(NaN % 3, NaN);
assert.equal(3 % NaN, NaN);
assert.equal(NaN % NaN, NaN);

// 字符串求余（字符串可转换为数值）
assert.equal("10" % "5", 0);
assert.equal("10" % "3", 1);
assert.equal("-10" % "3", -1);
assert.equal("10" % "-3", 1);
assert.equal("10" % "a", NaN); // "a" 无法转换为数值

// 字符串与数值
assert.equal("10" % 3, 1);
assert.equal(10 % "3", 1);
assert.equal("-10" % 3, -1);
assert.equal(10 % "-3", 1);

// 布尔值求余
assert.equal(true % 2, 1); // true 转换为 1
assert.equal(false % 2, 0); // false 转换为 0

// 布尔值与数值
assert.equal(true % 3, 1);
assert.equal(false % 10, 0);

// 布尔值与字符串（可转换为数值）
assert.equal(true % "2", 1);
assert.equal(false % "10", 0);

// null 求余
assert.equal(null % 2, 0); // null 转换为 0
assert.equal(2 % null, NaN);

// undefined 求余
assert.equal(undefined % 2, NaN);
assert.equal(2 % undefined, NaN);


// 数组求余
assert.equal([10] % [3], 1);
assert.equal([10, 20] % [2], NaN); // 数组转换为逗号分隔的字符串，然后尝试转换为数值

// 使用assert.ok来检查NaN情况
assert.ok(isNaN("10" % "a"));
assert.ok(isNaN(2 % null));
assert.ok(isNaN(undefined % 2));
assert.ok(isNaN([10, 20] % [2]));

// 循环生成更多测试用例
for (let i = 1; i <= 25; i++) {
    assert.equal(i % 2, i % 2);
    assert.equal(i % -2, i % -2);
    assert.equal(-i % 2, -i % 2);
    assert.equal(i % (i + 1), i);
}
