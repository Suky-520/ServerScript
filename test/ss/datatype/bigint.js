import assert from "assert";

console.log("测试语法::大整数");

// 整型
let a = 12n;
let b = 12345678901234567890123456789012345678901234567890n;
assert.ok(a === 12n);
assert.ok(b === 12345678901234567890123456789012345678901234567890n);

// 加法
for (let i = 0n; i <= 50n; i++) {
    assert.equal(i + i, 2n * i);
}

// 减法
for (let i = 0n; i <= 50n; i++) {
    assert.equal(i - i, 0n,);
}

// 乘法
for (let i = 1n; i <= 50n; i++) {
    assert.equal(i * i, i ** 2n,);
}

// 除法
for (let i = 1n; i <= 50n; i++) {
    assert.equal(i * 2n / i, 2n,);
}

// 取模
for (let i = 2n; i <= 50n; i++) {
    assert.equal((i * 2n + 1n) % i, 1n,);
}

// 比较运算
assert.equal(1n < 2n, true, '1n < 2n 应该为 true');
assert.equal(2n <= 2n, true, '2n <= 2n 应该为 true');
assert.equal(3n > 2n, true, '3n > 2n 应该为 true');
assert.equal(2n >= 2n, true, '2n >= 2n 应该为 true');
assert.equal(2n == 2n, true, '2n == 2n 应该为 true');
assert.equal(2n != 3n, true, '2n != 3n 应该为 true');

// 位运算 注意：BigInt位运算需要使用BigInt字面量
/*for (let i = 0n; i <= 50n; i++) {
    assert.equal(i & (i + 1n), i,);
    assert.equal(i | (i + 1n), i + 1n, );
    assert.equal(i ^ (i + 1n), 1n,);
}*/

// 左移和右移
for (let i = 0n; i <= 10n; i++) {
    assert.equal(i << 1n, i * 2n,);
    assert.equal(i >> 1n, i / 2n,);
}

// 请注意，BigInt运算中不允许除以0，尝试这样做会抛出一个RangeError。
// assert.throws(() => 1n / 0n, RangeError);

// 边界值测试
assert.equal((2n ** 53n) - 1n + 1n, 2n ** 53n, '2^53 - 1n + 1n 应该等于 2^53');
assert.equal(-(2n ** 53n) + 1n, -(2n ** 53n - 1n), '-(2^53) + 1n 应该等于 -(2^53 - 1n)');
assert.equal((2n ** 63n) - 1n + 1n, 2n ** 63n, '2^63 - 1n + 1n 应该等于 2^63');
assert.equal(-(2n ** 63n) + 1n, -(2n ** 63n - 1n), '-(2^63) + 1n 应该等于 -(2^63 - 1n)');

// 一元运算符
assert.equal(-(-1n), 1n, '-(-1n) 应该等于 1n');
assert.equal(+(+1n), 1n, '+(+1n) 应该等于 1n'); // 注意：一元加号不适用于BigInt

// 逻辑非
assert.equal(!0n, true, '!0n 应该为 true');
assert.equal(!1n, false, '!1n 应该为 false');

// 逻辑与
assert.equal((1n && 2n), 2n, '1n && 2n 应该等于 2n');

// 逻辑或
assert.equal((0n || 2n), 2n, '0n || 2n 应该等于 2n');

// 严格等于和严格不等于
assert.equal(1n === 1n, true, '1n === 1n 应该为 true');
assert.equal(1n !== 2n, true, '1n !== 2n 应该为 true');

// 更多位运算符
// 按位非
assert.equal(~1n, -2n, '~1n 应该等于 -2n');

// 左移和右移补充
for (let i = 0n; i <= 10n; i++) {
    assert.equal(i << 2n, i * 4n,);
    assert.equal(i >> 2n, i / 4n,);
}

// 无符号右移（BigInt不支持无符号右移，下面的测试用例仅作为演示，实际上会抛出异常）
// assert.throws(() => 1n >>> 1n, TypeError);

// 幂运算
assert.equal(2n ** 4n, 16n, '2n ** 4n 应该等于 16n');

// 综合测试
assert.equal((1n << 2n) + (5n & 3n), 5n, '(1n << 2n) + (5n & 3n) 应该等于 5n');
assert.equal((10n / 2n) * (4n | 1n), 25n, '(10n / 2n) * (4n | 1n) 应该等于 20n');

// 注意：以上测试中，部分操作如一元加号对BigInt是不适用的，逻辑非只对0n返回true，其他所有BigInt值都视为true。
// 无符号右移在BigInt中不支持，尝试使用会抛出TypeError异常。