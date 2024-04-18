import assert from "assert";

console.log("测试语法 条件三元运算符");

// 测试 condition 为真时，返回 expr1
assert.equal(true ? 1 : 0, 1);

// 测试 condition 为假时，返回 expr2
assert.equal(false ? 1 : 0, 0);

// 测试 condition 为 true 时，返回非基本类型的表达式
assert.deepEqual(true ? [1, 2, 3] : [], [1, 2, 3]);

// 测试 condition 为 false 时，返回非基本类型的表达式
assert.deepEqual(false ? [1, 2, 3] : [4, 5, 6], [4, 5, 6]);

// 测试 condition 为 true 时，返回字符串表达式
assert.equal(true ? 'yes' : 'no', 'yes');

// 测试 condition 为 false 时，返回字符串表达式
assert.equal(false ? 'yes' : 'no', 'no');

// 测试 condition 为真时，返回数值表达式
assert.equal(10 > 5 ? 10 : 5, 10);

// 测试 condition 为假时，返回数值表达式
assert.equal(5 > 10 ? 10 : 5, 5);

// 测试 condition 为 true 时，返回布尔表达式
assert.equal(1 === 1 ? true : false, true);

// 测试 condition 为 false 时，返回布尔表达式
assert.equal(1 === 2 ? true : false, false);
