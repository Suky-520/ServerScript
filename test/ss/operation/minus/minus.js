import assert from "assert";

console.log("测试运算符:: -");

// 数值之间的减法
assert.equal(10 - 5, 5);
assert.equal(-10 - 5, -15);
assert.equal(0.1 - 0.2, -0.1);

// 数值与字符串
assert.equal("10" - "5", 5);
assert.equal("10" - "test", NaN);
assert.equal("10.5" - "0.5", 10);

// 数值与布尔值
assert.equal(true - 1, 0); // true 转换为 1
assert.equal(false - 1, -1); // false 转换为 0

// 数值与 null
assert.equal(10 - null, 10); // null 转换为 0

// 数值与 undefined
assert.ok(isNaN(10 - undefined)); // undefined 转换为 NaN

// 字符串之间的减法
assert.equal("20" - "10", 10);
assert.ok(isNaN("hello" - "world"));

// 字符串与布尔值
assert.equal("10" - true, 9); // true 转换为 1
assert.equal("10" - false, 10); // false 转换为 0

// 字符串与 null
assert.equal("10" - null, 10); // null 转换为 0

// 字符串与 undefined
assert.ok(isNaN("10" - undefined)); // undefined 转换为 NaN

// 布尔值之间的减法
assert.equal(true - false, 1); // true 转换为 1，false 转换为 0

// 布尔值与 null
assert.equal(true - null, 1); // null 转换为 0

// 布尔值与 undefined
assert.ok(isNaN(false - undefined)); // undefined 转换为 NaN

// null 与 undefined
assert.ok(isNaN(null - undefined)); // null 转换为 0，undefined 转换为 NaN

// 数组与数值
assert.equal([10] - 5, 5); // 数组转换为其唯一元素的值 10
assert.ok(isNaN([1, 2] - 3)); // 数组转换为字符串 "1,2"，然后尝试与 3 进行减法

// 使用assert.ok来检查NaN情况，因为NaN !== NaN
assert.ok(isNaN("hello" - "world"));
assert.ok(isNaN("10" - undefined));
assert.ok(isNaN(false - undefined));
assert.ok(isNaN(null - undefined));
assert.ok(isNaN({toString: () => "5"} - "2"));
assert.ok(isNaN([1, 2] - 3));



