import assert from "assert";

console.log("测试运算符:: or");

// 布尔值逻辑或
assert.equal(true || false, true);
assert.equal(false || true, true);
assert.equal(false || false, false);

// 数值逻辑或
assert.equal(1 || 2, 1); // 第一个真值，返回第一个操作数
assert.equal(0 || 2, 2); // 第一个假值，返回第二个操作数
assert.equal(-1 || 2, -1); // 负数被视为真值

// 字符串逻辑或
assert.equal('hello' || 'world', 'hello'); // 非空字符串被视为真值
assert.equal('' || 'world', 'world'); // 空字符串被视为假值
assert.equal('hello' || '', 'hello'); // 非空字符串被视为真值

// null 和 undefined 逻辑或
assert.equal(null || 'hello', 'hello');
assert.equal('hello' || null, 'hello');
assert.equal(undefined || 'world', 'world');
assert.equal('world' || undefined, 'world');

// 数组和对象逻辑或
assert.deepEqual([] || 'hello', []); // 空数组被视为真值
assert.deepEqual({} || 'world', {}); // 空对象被视为真值

// 函数逻辑或
function myFunc() { return 'hello'; }
assert.equal(myFunc() || 'world', 'hello'); // 函数调用返回值与字符串

// 混合类型逻辑或
assert.equal(1 || 'hello', 1); // 数值和字符串
assert.deepEqual(false || {}, {}); // 布尔值和对象
assert.equal('yes' || [1, 2, 3], 'yes'); // 字符串和数组

// 特殊数值逻辑或
assert.equal(NaN || 'hello', 'hello'); // NaN 被视为假值
assert.equal('hello' || NaN, 'hello'); // 第一个操作数是真值
assert.equal(Infinity || 'world', Infinity); // Infinity 被视为真值
assert.equal('world' || Infinity, 'world'); // 第一个操作数是真值

// 使用assert.ok来检查布尔值结果
assert.ok(true || false);
assert.equal((false || false),false);
