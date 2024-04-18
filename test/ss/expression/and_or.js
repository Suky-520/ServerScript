import assert from "assert";

console.log("测试逻辑与或非");

// 基本类型
assert.equal(true && true, true); // 基本的AND运算
assert.equal(true && false, false);
assert.equal(false && true, false);
assert.equal(false && false, false);
assert.equal(true || true, true); // 基本的OR运算
assert.equal(true || false, true);
assert.equal(false || true, true);
assert.equal(false || false, false);
assert.equal(0 && 1, 0); // 使用数字进行AND运算
assert.equal(1 && 2, 2);
assert.equal(1 && 0, 0);
assert.equal(0 || 1, 1); // 使用数字进行OR运算
assert.equal(1 || 2, 1);
assert.equal(0 || 0, 0);
assert.equal("a" && "b", "b"); // 使用字符串进行AND运算
assert.equal("" && "b", "");
assert.equal("a" || "b", "a"); // 使用字符串进行OR运算
assert.equal("" || "b", "b");
assert.equal(null && undefined, null); // 使用null和undefined
assert.equal(null || undefined, undefined);
assert.equal((1 || 2) && (2 || 3), 2); // 带有括号的混合运算
assert.equal((1 && 2) || (2 && 3), 2);
assert.equal((false || true) && (0 || 1), 1);
assert.equal((true && false) || (null || "test"), "test");
assert.equal(1 && (null || "a" && "b"), "b"); // 更复杂的混合运算

// bigint
assert.equal(1n && 2n, 2n); // 使用BigInt进行AND运算
assert.equal(0n && 1n, 0n);
assert.equal(1n || 2n, 1n); // 使用BigInt进行OR运算
assert.equal(0n || 1n, 1n);
assert.equal(1n && 0n, 0n);
assert.equal((1n || 2n) && (2n || 3n), 2n); // 带有括号的BigInt混合运算
assert.equal((0n && 1n) || (2n && 3n), 3n);
assert.equal(1n && (0n || "a" && "b"), "b"); // BigInt与字符串的混合运算
assert.equal("a" && 1n, 1n); // 字符串与BigInt的AND运算
assert.equal("a" || 2n, "a"); // 字符串与BigInt的OR运算
assert.equal(0n || "b", "b");
assert.equal((1n && 2n) || ("test" && 3n), 2n);
assert.equal((0n || 1n) && ("a" || 2n), "a");
assert.equal((1n || 0n) && (2n || "b"), 2n);
assert.equal("a" && (1n || 2n), 1n); // 使用括号的字符串与BigInt混合运算


let counterA = 0;

function funcA() {
    counterA++;
    return false; // 或其他真值，取决于测试的需要
}

let counterB = 0;

function funcB() {
    counterB++;
    return true; // 或其他假值，取决于测试的需要
}

// 测试用例
// 测试&&短路行为，funcA返回false，预期不调用funcB
funcA() && funcB();
assert.equal(counterB, 0, "funcB should not be called due to && short-circuiting");

counterA = 0; // 重置计数器
counterB = 0;

// 测试||短路行为，funcA返回true，预期不调用funcB
funcA() || funcB();
assert.equal(counterB, 1, "funcB should not be called due to || short-circuiting");

counterA = 0; // 重置计数器
counterB = 0;

// 测试没有短路发生的情况
// 假设funcA()返回true，测试&&
funcA() && funcB();
assert.equal(counterB, 0, "funcB should be called because funcA returns true");

counterA = 0; // 重置计数器
counterB = 0;

// 假设funcA()返回false，测试||
funcA() || funcB();
assert.equal(counterB, 1, "funcB should be called because funcA returns false");

