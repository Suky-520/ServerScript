import assert from "assert";

console.log("测试基本数据类型::number");

// 测试Number.isFinite()的各种情况
assert.equal(Number.isFinite(Infinity), false, "Infinity不是有限数");
assert.equal(Number.isFinite(-Infinity), false, "-Infinity不是有限数");
assert.equal(Number.isFinite(NaN), false, "NaN不是有限数");
assert.equal(Number.isFinite('0'), false, "'0'不是Number类型，应返回false");
assert.equal(Number.isFinite(null), false, "null不是Number类型，应返回false");
assert.equal(Number.isFinite(0), true, "0是有限数");
assert.equal(Number.isFinite(2e64), true, "2e64是有限数");

// 测试Number.isInteger()的各种情况
assert.equal(Number.isInteger(0), true, "0是整数");
assert.equal(Number.isInteger(1), true, "1是整数");
assert.equal(Number.isInteger(-100000), true, "-100000是整数");
assert.equal(Number.isInteger(0.1), false, "0.1不是整数");
assert.equal(Number.isInteger(Math.PI), false, "π不是整数");
assert.equal(Number.isInteger(Infinity), false, "Infinity不是整数");
assert.equal(Number.isInteger(-Infinity), false, "-Infinity不是整数");
assert.equal(Number.isInteger("10"), false, "字符串'10'不是整数");
assert.equal(Number.isInteger(null), false, "null不是整数");
assert.equal(Number.isInteger(undefined), false, "undefined不是整数");
assert.equal(Number.isInteger(NaN), false, "NaN不是整数");

// 测试Number.isNaN()的各种情况
assert.equal(Number.isNaN(NaN), true, "NaN应该返回true");
assert.equal(Number.isNaN(0 / 0), true, "0除以0应该返回true");
assert.equal(Number.isNaN(NaN), true, "Number.NaN应该返回true");
assert.equal(Number.isNaN(0), false, "0不是NaN");
assert.equal(Number.isNaN(null), false, "null不是NaN");
assert.equal(Number.isNaN(undefined), false, "undefined不是NaN");
assert.equal(Number.isNaN('NaN'), false, "'NaN'字符串不是NaN");
assert.equal(Number.isNaN(''), false, "空字符串不是NaN");
assert.equal(Number.isNaN(true), false, "布尔值true不是NaN");

// 正常值
assert.equal(Number.isSafeInteger(0), true, "0是安全整数");
assert.equal(Number.isSafeInteger(1), true, "1是安全整数");
assert.equal(Number.isSafeInteger(-1), true, "-1是安全整数");
assert.equal(Number.isSafeInteger(Math.pow(2, 53) - 1), true, "2^53 - 1是安全整数的上限");
assert.equal(Number.isSafeInteger(-(Math.pow(2, 53) - 1)), true, "-(2^53 - 1)是安全整数的下限");

// 边界值
assert.equal(Number.isSafeInteger(Math.pow(2, 53)), false, "2^53超出安全整数范围");
assert.equal(Number.isSafeInteger(-(Math.pow(2, 53))), false, "-2^53超出安全整数范围");

// 非法值
assert.equal(Number.isSafeInteger("1024"), false, "字符串不是安全整数");
assert.equal(Number.isSafeInteger(NaN), false, "NaN不是安全整数");
assert.equal(Number.isSafeInteger(Infinity), false, "Infinity不是安全整数");
assert.equal(Number.isSafeInteger(-Infinity), false, "-Infinity不是安全整数");
assert.equal(Number.isSafeInteger(3.14), false, "浮点数不是安全整数");

// 正常值
assert.equal(Number.parseFloat("3.14"), 3.14, "'3.14' 应解析为 3.14");
assert.equal(Number.parseFloat("10.00"), 10.00, "'10.00' 应解析为 10");

// 边界值和非法值
assert.equal(Number.parseFloat("34 45 66"), 34, "'34 45 66' 应解析为 34，因为解析停在第一个空格");
assert.equal(Number.parseFloat("   60 "), 60, "'   60 ' 应解析为 60，忽略前后空格");

// 特殊情况
assert.equal(Number.parseFloat(""), NaN, "空字符串应解析为 NaN");
assert.equal(Number.parseFloat("abc"), NaN, "'abc' 应解析为 NaN");
assert.equal(Number.parseFloat("NaN"), NaN, "'NaN' 字符串应解析为 NaN");
assert.equal(Number.parseFloat("123abc"), 123, "'123abc' 应解析为 123，忽略数字后的文本");

// 确保使用assert.strictEqual来比较NaN值
assert.equal(Number.isNaN(Number.parseFloat("abc")), true, "'abc' 应解析为 NaN");

// 正常值
assert.equal(Number.parseInt('0x10'), 16, "'0x10' 应该被解析为十六进制整数 16");
assert.equal(Number.parseInt('10'), 10, "'10' 应该被解析为整数 10");
assert.equal(Number.parseInt('010'), 10, "'010' 应该被解析为整数 10，忽略前导零");

// 非法值和边界值
assert.equal(Number.parseInt(''), NaN, "空字符串应该被解析为 NaN");
assert.equal(Number.parseInt('abc'), NaN, "'abc' 应该被解析为 NaN");
assert.equal(Number.parseInt('10abc'), 10, "'10abc' 应该被解析为整数 10，忽略后续非数字字符");

// 使用assert.equal检查NaN
assert.equal(Number.isNaN(Number.parseInt('abc')), true, "'abc' 解析为 NaN");

// 正常值
assert.equal((123.456).toFixed(2), '123.46', '两位小数的四舍五入处理失败');
assert.equal((0).toFixed(0), '0', '整数0转换失败');
assert.equal((-123.456).toFixed(3), '-123.456', '负数保留三位小数失败');

// 边界值
assert.equal((1.005).toFixed(2), '1.00', '边界值四舍五入失败');
assert.equal((Number.MIN_VALUE).toFixed(2), '0.00', '最小值转换失败');

// 正常值测试
assert.equal((123.456).toPrecision(4), '123.5', '四舍五入失败');
assert.equal((123).toPrecision(3), '123', '整数返回失败');
assert.equal((0.000123).toPrecision(2), '0.00012', '小数点位数错误');
assert.equal((100000000000000).toPrecision(3), '1.00e+14', '科学计数法错误');

// 边界值测试
assert.equal((0.000000000000001).toPrecision(2), '1.0e-15', '极小值转换失败');
assert.equal((100000000000000000000).toPrecision(2), '1.0e+20', '极大值转换失败');
assert.equal((1.0001).toPrecision(3), '1.00', '超过小数点位数的四舍五入错误');

// 正常值测试
assert.equal((123).toString(), '123', '整数转换失败');
assert.equal((123.456).toString(), '123.456', '小数转换失败');
assert.equal((0).toString(), '0', '零转换失败');
assert.equal((0.1 + 0.2).toString(), '0.30000000000000004', '小数运算结果转换失败');

// 边界值测试
assert.equal((Number.MAX_VALUE).toString(), Number.MAX_VALUE.toString(), '最大值转换失败');
assert.equal((Number.MIN_VALUE).toString(), Number.MIN_VALUE.toString(), '最小值转换失败');

// 非法值测试
assert.equal((NaN).toString(), 'NaN', 'NaN值转换失败');
assert.equal((Infinity).toString(), 'Infinity', '正无穷大值转换失败');
assert.equal((-Infinity).toString(), '-Infinity', '负无穷大值转换失败');

// 正常值测试
assert.equal((123).toString(2), '1111011', '二进制转换失败');
assert.equal((123).toString(8), '173', '八进制转换失败');
assert.equal((123).toString(16), '7b', '十六进制转换失败');
assert.equal((123).toString(10), '123', '十进制转换失败');

// 边界值测试
assert.equal((0).toString(2), '0', '零的二进制转换失败');
assert.equal((0).toString(8), '0', '零的八进制转换失败');
assert.equal((0).toString(16), '0', '零的十六进制转换失败');
assert.equal((0).toString(10), '0', '零的十进制转换失败');

assert.equal((255).toString(2), '11111111', '255的二进制转换失败');
assert.equal((255).toString(8), '377', '255的八进制转换失败');
assert.equal((255).toString(16), 'ff', '255的十六进制转换失败');
assert.equal((255).toString(10), '255', '255的十进制转换失败');

// 非法值测试
assert.equal((NaN).toString(2), 'NaN', 'NaN值转换失败');
assert.equal((Infinity).toString(2), 'Infinity', '正无穷大值转换失败');
assert.equal((-Infinity).toString(2), '-Infinity', '负无穷大值转换失败');

// 正常值测试
assert.equal((123.456).toExponential(2), '1.23e+02', '正常小数转换失败');
assert.equal((123456).toExponential(2), '1.23e+05', '正常整数转换失败');
assert.equal((0.000123).toExponential(2), '1.23e-04', '小于1的小数转换失败');
assert.equal((0.1 + 0.2).toExponential(2), '3.00e-01', '小数运算结果转换失败');

// 边界值测试
assert.equal((Number.MAX_VALUE).toExponential(2), '1.80e+308', '最大值转换失败');
assert.equal((Number.MIN_VALUE).toExponential(2), '4.94e-324', '最小值转换失败');

// 非法值测试
assert.equal((NaN).toExponential(2), 'NaN', 'NaN值转换失败');
assert.equal((Infinity).toExponential(2), 'Infinity', '正无穷大值转换失败');
assert.equal((-Infinity).toExponential(2), '-Infinity', '负无穷大值转换失败');