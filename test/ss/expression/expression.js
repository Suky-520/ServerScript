import assert from "assert";

console.log("测试表达式");

// 使用后缀自增和自减运算符
let a = 1;
assert.equal(a++, 1); // a自增前的值是1
assert.equal(++a, 3); // a再自增后的值是3
assert.equal(a--, 3); // a自减前的值是3
assert.equal(--a, 1); // a再自减后的值是1

// 取反操作
let b = 1;
assert.equal(-b, -1); // 取反操作
b = -1;
assert.equal(-b, 1); // 再次取反

// 与或非运算
let c = false;
assert.equal(!c, true); // 逻辑非操作
assert.equal(!false, true);
assert.equal(!0, true);
assert.equal(!1, false);
assert.equal(2 ** 3, 8); // 指数运算
assert.equal(2 * 3, 6); // 乘法
assert.equal(4 / 2, 2); // 除法
assert.equal(4 % 3, 1); // 取模
assert.equal(3 + 2, 5); // 加法
assert.equal(3 - 2, 1); // 减法
assert.equal(3 >> 1, 1); // 右移
assert.equal(3 << 1, 6); // 左移
assert.equal(8 >>> 2, 2); // 无符号右移
assert.ok(8 == 8); // 等于
assert.ok(8 != "2"); // 不等于
assert.ok(8 === 8); // 全等于
assert.ok(8 !== "8"); // 不全等于
assert.equal(4 & 5, 4); // 按位与
assert.equal(4 ^ 5, 1); // 按位异或
assert.equal(4 | 5, 5); // 按位或
assert.equal(4 == 4 && 5 == 4, false); // 逻辑与
assert.equal(4 == 4 && 1 == 1, true);
assert.equal(4 == 4 || 5 == 4, true); // 逻辑或
assert.equal(4 == 4 || 1 == 1, true);
assert.equal(4 > 0 ? true : false, true); // 三元运算符

// 运算赋值操作
let e = 2;
e *= 2;
assert.equal(e, 4);
e /= 2;
assert.equal(e, 2);
e %= 1;
assert.equal(e, 0);
e += 1;
assert.equal(e, 1);
e -= 1;
assert.equal(e, 0);
e <<= 2;
assert.equal(e, 0);
e >>= 2;
assert.equal(e, 0);
e >>>= 2;
assert.equal(e, 0);
e = 4;
e &= 5;
assert.equal(e, 4);
e = 4;
e ^= 5;
assert.equal(e, 1);
e = 4;
e |= 5;
assert.equal(e, 5);
e = 2;
e **= 4;
assert.equal(e, 16);

// 字符串连接
assert.equal("Hello, " + "world!", "Hello, world!");

// 模板字符串
let name = "world";
assert.equal(`Hello, ${name}!`, "Hello, world!");

// typeof 函数
assert.equal(typeOf("string").typeString, "string");
assert.equal(typeOf(1).typeString, "number");

// in 运算符
let myObject = {a: 1};
assert.ok("a" in myObject);

// null 和 undefined 的比较
assert.equal(null == undefined, true);
assert.equal(null === undefined, false);

// delete 运算符
let myObj = {property: true};
delete myObj.property;
assert.equal(myObj.property, undefined);

// 以下部分是对!运算符不同值的测试
assert.equal(!undefined, true);     // undefined
assert.equal(!null, true);          // null
assert.equal(!true, false);         // 布尔值
assert.equal(!false, true);         // 布尔值
assert.equal(!1, false);            // 任何非零数字
assert.equal(!0, true);             // 0
assert.equal(!-1, false);           // 负数也是真值
assert.equal(!NaN, true);           // NaN 被视为假值
assert.equal(!"", true);            // 空字符串是假值
assert.equal(!"some text", false);  // 任何非空字符串是真值
assert.equal(!{}, false);           // 任何对象（包括空对象）是真值
assert.equal(![], false);           // 任何数组（包括空数组）是真值
assert.equal(!-0, true);            // 特殊案例：-0
assert.equal(!0n, true);            // BigInt(0) 是假值
assert.equal(!1n, false);           // 非零 BigInt 是真值
function testFunction() {
}          // 函数
assert.equal(!testFunction, false); // 函数是真值

// 加法
assert.equal(10 + 5, 15);

// 减法
assert.equal(10 - 5, 5);

// 乘法
assert.equal(10 * 5, 50);

// 除法
assert.equal(10 / 5, 2);

// 取模
assert.equal(10 % 3, 1);

// 指数
assert.equal(2 ** 3, 8);

// 直接赋值
a = 10;
assert.equal(a, 10);

// 加法赋值
a += 5;
assert.equal(a, 15);

// 减法赋值
a -= 5;
assert.equal(a, 10);

// 乘法赋值
a *= 2;
assert.equal(a, 20);

// 除法赋值
a /= 2;
assert.equal(a, 10);

// 取模赋值
a %= 3;
assert.equal(a, 1);

// 指数赋值
a **= 2;
assert.equal(a, 1);

// 左移赋值
a = 1; // Reset
a <<= 2;
assert.equal(a, 4);

// 右移赋值
a >>= 1;
assert.equal(a, 2);

// 无符号右移赋值
a >>>= 1;
assert.equal(a, 1);

// 按位与赋值
a &= 3;
assert.equal(a, 1);

// 按位异或赋值
a ^= 3;
assert.equal(a, 2);

// 按位或赋值
a |= 2;
assert.equal(a, 2);

// 等于
assert.equal(10 == "10", true);
assert.equal(true == 1, true);
assert.equal(false == 0, true);
assert.equal(true == "1", true);
assert.equal(false == "", true);
assert.equal(null == undefined, true);
assert.equal(NaN == NaN, false);
assert.equal(10 == 10, true);
assert.equal("hello" == "hello", true);
assert.equal("0" == 0, true);

// 全等于
assert.equal(10 === "10", false);
assert.equal(true === 1, false);
assert.equal(0 === false, false);
assert.equal(null === undefined, false);
assert.equal(NaN === NaN, false);
assert.equal(10 === 10, true);
assert.equal("hello" === "hello", true);
let obj = {};
assert.equal(obj === obj, true);
assert.equal({key: "value"} === {key: "value"}, false);
let arr = [1, 2, 3];
assert.equal(arr === arr, true);
assert.equal([1, 2, 3] === [1, 2, 3], false);
assert.equal(true === true, true);
assert.equal(false === false, true);

// 不等于
assert.equal(10 != "10", false);
assert.equal(10 != "10", false);
assert.equal(true != 1, false);
assert.equal(false != 0, false);
assert.equal(true != "1", false);
assert.equal(false != "", false);
assert.equal(null != undefined, false);
assert.equal(NaN != NaN, true);
assert.equal(10 != 10, false);
assert.equal("hello" != "hello", false);
assert.equal("0" != 0, false);

// 不全等于
assert.equal(10 !== "10", true);
assert.equal("5" !== 5, true);
assert.equal(true !== 1, true);
assert.equal(false !== "false", true);
assert.equal(null !== undefined, true);
assert.equal(NaN !== NaN, true);
assert.equal(10 !== 5, true);
assert.equal("hello" !== "world", true);
assert.equal(10 !== 10, false);
assert.equal("hello" !== "hello", false);
let obj1 = {key: "value"};
let obj2 = {key: "value"};
let obj3 = obj1;
assert.equal(obj1 !== obj2, true); // 不同对象，即使内容相同
assert.equal(obj1 !== obj3, false); // 相同的引用
assert.equal([1, 2, 3] !== [1, 2, 3], true); // 不同数组，即使内容相同

// 大于
assert.equal(10 > 5, true);
assert.equal(5 > 10, false);
assert.equal(10 > 10, false);
assert.equal("banana" > "apple", true);
assert.equal("hello" > "hello", false);
assert.equal(true > false, true);
assert.equal(false > true, false);
assert.equal(null > 0, false);
assert.equal(1 > null, true);
assert.equal(undefined > 0, false);
assert.equal(0 > undefined, false);
assert.equal(NaN > NaN, false);
assert.equal(NaN > 0, false);
assert.equal(10 > "5", true);
assert.equal(5 > "10", false);
assert.equal(5 > "hello", false); // "hello"转换为NaN
assert.equal(2 > true, true);
assert.equal(0 > false, false);

// 小于
assert.equal(10 < 15, true);
assert.equal(5 < 10, true);
assert.equal(20 < 15, false);
assert.equal(10 < 10, false);
assert.equal("apple" < "banana", true);
assert.equal("hello" < "hello", false);
assert.equal("banana" < "apple", false);
assert.equal(false < true, true);
assert.equal(true < false, false);
assert.equal(null < 1, true);
assert.equal(0 < null, false);
assert.equal(undefined < 0, false);
assert.equal(0 < undefined, false);
assert.equal(NaN < NaN, false);
assert.equal(NaN < 0, false);
assert.equal("5" < 10, true);
assert.equal("10" < 5, false);
assert.equal("hello" < 5, false); // "hello"转换为NaN
assert.equal(true < 2, true);
assert.equal(false < 0, false);

// 大于等于
assert.equal(5 >= 5, true);
assert.equal(10 >= 5, true);
assert.equal(3 >= 5, false);
assert.equal("apple" >= "banana", false);
assert.equal("hello" >= "hello", true);
assert.equal(true >= false, true);
assert.equal(false >= true, false);
assert.equal(null >= 0, true);
assert.equal(null >= 1, false);
assert.equal(undefined >= 0, false);
assert.equal(0 >= undefined, false);
assert.equal(NaN >= NaN, false);
assert.equal(NaN >= 0, false);
assert.equal("5" >= 5, true);
assert.equal("10" >= 5, true);
assert.equal("1" >= 5, false); // "hello"转换为NaN
assert.equal(true >= 1, true);
assert.equal(false >= 1, false);

// 小于等于
assert.equal(10 <= 10, true);
assert.equal(5 <= 10, true);
assert.equal(15 <= 10, false);
assert.equal('abc' <= 'abc', true);
assert.equal('abc' <= 'bcd', true);
assert.equal('bcd' <= 'abc', false);
assert.equal(5 <= '10', true);
assert.equal('10' <= 10, true);
assert.equal(15 <= '10', false);
assert.equal(NaN <= 10, false);
assert.equal(10 <= NaN, false);
assert.equal(Infinity <= Infinity, true);
assert.equal(-Infinity <= Infinity, true);
assert.equal(Infinity <= -Infinity, false);
assert.equal(undefined <= 10, false);
assert.equal(10 <= undefined, false);
assert.equal(null <= 0, true);
assert.equal(0 <= null, true);
assert.equal(true <= 1, true);
assert.equal(false <= 0, true);
assert.equal(1 <= true, true);
assert.equal(0 <= false, true);
assert.deepEqual(10 <= {}, false);

// 逻辑与
assert.equal(true && true, true);
assert.equal(true && false, false);
assert.equal(false && true, false);
assert.equal(false && false, false);
assert.equal(null && undefined, null);
assert.equal(undefined && null, undefined);
assert.equal(1 && 0, 0);
assert.equal(0 && 1, 0);
assert.equal(-1 && 0, 0);
assert.equal(0 && -1, 0);
assert.equal(0 && NaN, 0);
assert.equal("" && "hello", "");
assert.equal("hello" && "", "");
assert.equal("hello" && "world", "world");
assert.deepEqual({} && null, null);
assert.deepEqual(null && {}, null);
assert.deepEqual([] && false, false);
assert.deepEqual(false && [], false);
assert.equal(Infinity && true, true);
assert.equal(false && Infinity, false);

// 逻辑或
assert.equal(true || true, true);
assert.equal(true || false, true);
assert.equal(false || true, true);
assert.equal(false || false, false);
assert.equal(null || undefined, undefined);
assert.equal(undefined || null, null);
assert.equal(1 || 0, 1);
assert.equal(0 || 1, 1);
assert.equal(-1 || 0, -1);
assert.equal(0 || -1, -1);
assert.equal(NaN || 0, 0);
assert.equal("" || "hello", "hello");
assert.equal("hello" || "", "hello");
assert.equal("hello" || "world", "hello");
assert.deepEqual({} || null, {});
assert.deepEqual(null || {}, {});
assert.deepEqual([] || false, []);
assert.deepEqual(false || [], []);
assert.equal(Infinity || true, Infinity);
assert.equal(false || Infinity, Infinity);

// 逻辑非
assert.equal(!true, false);
assert.equal(!false, true);
assert.equal(!null, true);
assert.equal(!undefined, true);
assert.equal(!1, false);
assert.equal(!0, true);
assert.equal(!-1, false);
assert.equal(!NaN, true);
assert.equal(!"", true);
assert.equal(!"hello", false);
assert.equal(!{}, false);
assert.equal(![], false);
assert.equal(!Infinity, false);
assert.equal(!0n, true);
assert.equal(!1n, false);
assert.equal(!(() => {
}), false);

// 按位与
assert.equal(5 & 1, 1);
// 按位或
assert.equal(5 | 2, 7);
// 任何数与0进行按位或操作，结果为其自身
assert.equal(4 | 0, 4)
// 任何数与自身进行按位或操作，结果为其自身
assert.equal(6 | 6, 6)
// 假设Number.MAX_SAFE_INTEGER代表Go中int64能表示的最大正数
assert.equal(9223372036854775807 | 0, 9223372036854775807)
// -1 (1111111111111111111111111111111111111111111111111111111111111111) | 2 (10) = -1
assert.equal(-1 | 2, -1)
// 2^62 (1000000000000000000000000000000000000000000000000000000000000000) | 1 (1) = 2^62 + 1
assert.equal((1 << 62) | 1, (1 << 62) + 1)
// 按位异或
assert.equal(5 ^ 1, 4);
// 基本测试用例
assert.equal(5 ^ 5, 0) // 任何数与自身异或结果为0
assert.equal(5 ^ 0, 5) // 任何数与0异或结果为其自身
assert.equal(5 ^ 1, 4) // 示例中给出的测试
// 边界值和特殊值测试用例
assert.equal(9223372036854775807 ^ -9223372036854775808, -1) // 最大正数与最小负数异或
assert.equal(-1 ^ -1, 0) // 负一与自身异或
assert.equal(9223372036854775807 ^ 0, 9223372036854775807) // 最大正数与0异或
//assert.equal(-9223372036854775808 ^ 0, -9223372036854775808) // 最小负数与0异或
assert.equal(9223372036854775807 ^ 9223372036854775807, 0) // 最大正数与自身异或
assert.equal(-9223372036854775808 ^ -9223372036854775808, 0) // 最小负数与自身异或

// 位移操作
assert.equal(5 << 1, 10);// 左移
assert.equal(-5 << 1, -10); // -5左移1位得到-10
assert.equal(0 << 1, 0); // 0左移任何位数仍然是0
assert.equal((2147483647 >> 1) << 1, 2147483646); // 使用右移然后左移回来，以避免直接左移导致的溢出
assert.equal((1 << 30) << 1, -2147483648); //
assert.equal((1 << 31) << 1, 0); // 超过32位整数范围的左移操作，导致溢出
assert.equal(-5 >> 1, -3); // 有符号右移

// 三元运算符
let age = 20;
assert.equal(age >= 18 ? "adult" : "minor", "adult");

// 按位非
assert.equal(~5, -6);
assert.equal(~-3, 2); // 对于64位，结果不变，因为-3在32位和64位下表现一致
assert.equal(~-1, 0); // -1的按位非是0
assert.equal(~0, -1); // 0的按位非是-1
assert.equal(~2147483647, -2147483648); // 32位整数的最大值
assert.equal(~-2147483648, 2147483647); // 32位整数的最小值
assert.equal(~4294967295, -4294967296); // 32位无符号整数的最大值
assert.equal(~-4294967296, 4294967295); // 对应于32位无符号整数的最大值的负数
assert.equal(~1, -2); // 测试1的按位非
assert.equal(~-2, 1); // 测试-2的按位非
assert.equal(~123456789, -123456790); // 一个较大的正数

//无符号右移
assert.equal(-5 >>> 32, 4294967291);
assert.equal(-1 >>> 0, 4294967295);
assert.equal(-5 >>> 1, 2147483645);
assert.equal(4294967296 >>> 32, 0);
assert.equal(-5 >>> 128, 4294967291);
assert.equal(0 >>> 32, 0);
assert.equal(-9223372036854775808 >>> 32, 0);
assert.equal(-5 >>> -32, 4294967291);

// 基本算术运算
assert.equal(1 + 2, 3, '1 + 2 应该等于 3');
assert.equal(2 * 3, 6, '2 * 3 应该等于 6');
assert.equal(4 - 2, 2, '4 - 2 应该等于 2');
assert.equal(8 / 2, 4, '8 / 2 应该等于 4');

// 字符串连接
assert.equal('Hello, ' + 'world!', 'Hello, world!', '字符串应该被正确连接');

// 类型转换
assert.equal('5' + 2, '52', '字符串和数字相加应该进行字符串连接');
assert.equal('5' - 2, 3, '字符串和数字相减应该转换字符串为数字进行计算');
assert.equal('5' * '2', 10, '两个字符串包含的数字相乘应该转换为数字进行计算');
assert.equal('10' / '2', 5, '两个字符串包含的数字相除应该转换为数字进行计算');

// 布尔逻辑
assert.equal(true && false, false, 'true AND false 应该是 false');
assert.equal(true || false, true, 'true OR false 应该是 true');
assert.equal(!true, false, 'NOT true 应该是 false');

// 对象属性访问
const obj4 = { key: 'value' };
assert.equal(obj4.key, 'value', '对象的属性应该可以被访问');

// 数组索引访问
const arr2 = [1, 2, 3];
assert.equal(arr2[0], 1, '数组的第一个元素应该是 1');
assert.equal(arr2.length, 3, '数组的长度应该是 3');

// 函数调用
function add(a, b) {
    return a + b;
}
assert.equal(add(1, 2), 3, '函数 add(1, 2) 应该返回 3');

// 类型转换为布尔值
assert.equal(!!'non-empty-string', true, '非空字符串转换为布尔值应该是 true');
assert.equal(!!'', false, '空字符串转换为布尔值应该是 false');
assert.equal(!!0, false, '数字 0 转换为布尔值应该是 false');
assert.equal(!!1, true, '非零数字转换为布尔值应该是 true');

// Null 和 Undefined
assert.equal(null == undefined, true, 'null 和 undefined 在非严格比较下应该相等');
assert.equal(null === undefined, false, 'null 和 undefined 在严格比较下不应该相等');
