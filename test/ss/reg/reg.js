import assert from "assert";

console.log("测试语法:: reg");

// 测试邮箱地址的正则表达式
let emailRegex = /^[^\s@]+@[^\s@]+\.[^\s@]+$/;
assert.ok(emailRegex.test('example@example.com'));
assert.ok(!emailRegex.test('example.com'));
assert.ok(!emailRegex.test('example@example'));

// 测试数字的正则表达式
let digitRegex = /^\d+$/;
assert.ok(digitRegex.test('12345'));
assert.ok(!digitRegex.test('12345abc'));
assert.ok(!digitRegex.test('abc'));

// 测试电话号码的正则表达式 (简单示例，可能不适用于所有格式)
let phoneRegex = /^\+?[0-9]{10,12}$/;
assert.ok(phoneRegex.test('+1234567890'));
assert.ok(phoneRegex.test('1234567890'));
assert.ok(!phoneRegex.test('123-456-7890'));

// 测试含有特定单词的正则表达式
let wordRegex = /\bhello\b/;
assert.ok(wordRegex.test('hello world'));
assert.ok(!wordRegex.test('helloworld'));
assert.ok(!wordRegex.test('Hello world'), "Case sensitive");

// 测试URL的正则表达式 (简化版)
let urlRegex = /^(https?:\/\/)?([\da-z\.-]+)\.([a-z\.]{2,6})([\/\w \.-]*)*\/?$/;
assert.ok(urlRegex.test('http://example.com'));
assert.ok(urlRegex.test('https://www.example.com/path/to/file'));
assert.ok(urlRegex.test('www.example.com'));

// 测试日期格式 YYYY-MM-DD 的正则表达式
let dateRegex = /^\d{4}-\d{2}-\d{2}$/;
assert.ok(dateRegex.test('2021-03-15'));
assert.ok(!dateRegex.test('15-03-2021'));
assert.ok(!dateRegex.test('03/15/2021'));

// 测试密码强度的正则表达式 (至少8个字符，包含字母和数字)
let passwordRegex = /^(?=.*[A-Za-z])(?=.*\d)[A-Za-z\d]{8,}$/;
assert.ok(passwordRegex.test('example123'));
assert.ok(!passwordRegex.test('example'));
assert.ok(!passwordRegex.test('12345678'));

// 使用assert.equal来验证测试结果是否符合预期
assert.equal(emailRegex.test('user@test.com'), true);
assert.equal(digitRegex.test('123abc'), false);
