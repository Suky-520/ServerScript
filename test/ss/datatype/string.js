import assert from "assert";

console.log("测试基本数据类型::string");

// 声明字符串变量
let a = "123abc";
// 单引号
let b = '123abc';
// 反引号
let c = `123abc`;
// 模板字符串
let d = `abc${c}`; //结果是:abc123abc
// 声明整数常量
const e = "123abc";
// 单引号
const f = '123abc';
// 反引号
const g = `123abc`;
// 模板字符串
let h = `abc${g}`; //结果是:abc123abc
assert.ok(a === "123abc");
assert.ok(b === "123abc");
assert.ok(c === "123abc");
assert.ok(d === "abc123abc");
assert.ok(e === "123abc");
assert.ok(f === "123abc");
assert.ok(g === "123abc");
assert.ok(h === "abc123abc");

// 基本表达式嵌入
assert.equal(`你好${1 + 54}`, '你好55');

// 多行字符串
assert.equal(`第一行
第二行`, '第一行\n第二行');

// 嵌入变量
let name = '世界';
assert.equal(`你好，${name}！`, '你好，世界！');

// 嵌入对象属性
let person = {firstName: '张', lastName: '三'};
assert.equal(`尊敬的${person.firstName}${person.lastName}`, '尊敬的张三');

// 嵌入三元运算符
let isMorning = true;
assert.equal(`现在是${isMorning ? '上午' : '下午'}`, '现在是上午');

// 嵌入表达式计算
let aa = 10, bb = 20;
assert.equal(`结果是${aa + bb}`, '结果是30');

// 嵌入布尔值
let condition = true;
assert.equal(`状态：${condition}`, '状态：true');

// 使用assert.ok来验证模板字符串生成的字符串是否符合预期
assert.ok(`你好${1 + 54}` === '你好55');

// 嵌入函数调用
assert.equal(`${greet('世界')}`, '你好，世界！');

function greet(name) {
    return `你好，${name}！`;
}

// 测试空字符串
assert.equal("".length, 0);

// 测试正常字符串
assert.equal("hello".length, 5);

// 测试正常情况
assert.equal("hello".charAt(1), 'e');

// 测试越界情况
assert.equal("hello".charAt(5), '');

//测试负值情况
assert.equal("hello".charAt(-1), '');

// 测试正常情况
assert.equal("hello".charCodeAt(1), 101);

// 测试越界情况
assert.equal("hello".charCodeAt(5), NaN);

// 测试负值情况
assert.equal("hello".charCodeAt(-1), NaN);

// 测试正常连接
assert.equal("hello".concat(" ", "world"), "hello world");

// 测试连接空字符串
assert.equal("hello".concat(""), "hello");

// 测试小写转大写
assert.equal("hello".toUpperCase(), "HELLO");

// 测试空字符串
assert.equal("".toUpperCase(), "");

// 测试大写转小写
assert.equal("HELLO".toLowerCase(), "hello");

// 测试空字符串
assert.equal("".toLowerCase(), "");

// 测试包含子字符串
assert.equal("hello world".includes("world"), true);

// 测试不包含子字符串
assert.equal("hello world".includes("abc"), false);

// 测试空字符串
assert.equal("hello world".includes(""), true);

// 测试找到子字符串
assert.equal("hello world".indexOf("world"), 6);

// 测试找不到子字符串
assert.equal("hello world".indexOf("abc"), -1);

// 测试空字符串
assert.equal("hello world".indexOf(""), 0);

// 测试找到子字符串
assert.equal("hello world world".lastIndexOf("world"), 12);

// 测试找不到子字符串
assert.equal("hello world".lastIndexOf("abc"), -1);

// 测试空字符串
assert.equal("hello world".lastIndexOf(""), 11);

// 正则匹配
assert.equal("hello world".match(/[a-z]+/)[0], "hello");

// 无匹配
assert.equal("hello world".match(/xyz/), null);

// 替换字符串
assert.equal("hello world".replace("world", "everyone"), "hello everyone");

// 使用正则表达式替换
assert.equal("hello world".replace(/\bworld\b/, "everyone"), "hello everyone");

// 无替换发生
assert.equal("hello world".replace("xyz", "everyone"), "hello world");

// 查找位置
assert.equal("hello world".search("world"), 6);

// 未找到
assert.equal("hello world".search("xyz"), -1);

// 正常截取
assert.equal("hello world".slice(0, 5), "hello");

// 越界截取
assert.equal("hello world".slice(0, 500), "hello world");

// 负值截取
assert.equal("hello world".slice(-5), "world");

// 正常分割
assert.equal("hello world".split(" ")[0], "hello");

// 分割符不存在
assert.equal("hello world".split("xyz").length, 1);

// 空字符串分割
assert.equal("hello".split("").length, 5);

// 正常截取
assert.equal("hello world".substring(0, 5), "hello");

// 越界截取
assert.equal("hello world".substring(0, 500), "hello world");

// 起始大于终止
assert.equal("hello world".substring(5, 0), "hello");

// 移除空格
assert.equal(" hello world ".trim(), "hello world");

// 不需要移除
assert.equal("hello world".trim(), "hello world");

// 移除开头空格
assert.equal(" hello world".trimStart(), "hello world");

// 不需要移除
assert.equal("hello world".trimStart(), "hello world");

// 移除末尾空格
assert.equal("hello world ".trimEnd(), "hello world");

// 不需要移除
assert.equal("hello world".trimEnd(), "hello world");

// 补全字符
assert.equal("hello".padStart(10, "x"), "xxxxxhello");

// 不需要补全
assert.equal("hello".padStart(5, "x"), "hello");

// 补全字符
assert.equal("hello".padEnd(10, "x"), "helloxxxxx");

// 不需要补全
assert.equal("hello".padEnd(5, "x"), "hello");

// 字符串比较(正序)
assert.equal("hello".localeCompare("world"), -1);

// 字符串比较(反序)
assert.equal("world".localeCompare("hello"), 1);

// 字符串比较(相同)
assert.equal("hello".localeCompare("hello"), 0);
