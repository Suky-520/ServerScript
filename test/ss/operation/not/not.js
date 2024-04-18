import assert from "assert"

console.log("测试运算符:: !");

// 布尔值
assert.equal(!true, false);
assert.equal(!false, true);

// 数值
assert.equal(!0, true); // 0 被视为假值
assert.equal(!1, false); // 非零数值被视为真值
assert.equal(!-1, false); // 负数被视为真值
assert.equal(!NaN, true); // NaN 被视为假值

// 字符串
assert.equal(!'', true); // 空字符串被视为假值
assert.equal(!'hello', false); // 非空字符串被视为真值

// null 和 undefined
assert.equal(!null, true); // null 被视为假值
assert.equal(!undefined, true); // undefined 被视为假值

// 对象和数组
assert.equal(!{}, false); // 对象被视为真值
assert.equal(![], false); // 数组被视为真值

// 函数
function myFunc() {}
assert.equal(!myFunc, false); // 函数被视为真值

// 特殊数值
assert.equal(!Infinity, false); // Infinity 被视为真值

// 使用assert.ok来检查布尔值结果
assert.ok(!false); // 验证真值
assert.ok(!0); // 验证真值
assert.ok(!''); // 验证真值
assert.ok(!null); // 验证真值
assert.ok(!undefined); // 验证真值
assert.ok(!NaN); // 验证真值

// 复合逻辑表达式
assert.equal(!!true, true); // 双重否定返回原布尔值
assert.equal(!!0, false); // 双重否定保持数值的布尔转换
