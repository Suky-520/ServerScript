import assert from "assert";

console.log("测试语法 fori");

function mes(i) {
    return `language/for/1.js 测试用例[${i}]`
}

let k = 1;
for (let i = 0; i < 10; i++) {
    k = i;
}
assert.equal(k, 9, mes(1));

for (k = 10; k > 0; k--) {

}
assert.equal(k, 0, mes(2));


for (let i = 0; i < 10; i++) {
    if (i > 5) {
        continue;
    }
    k = i;
}
assert.equal(k, 5, mes(3));

for (let i = 0; i < 10; i++) {
    k = i;
    if (i > 5) {
        break;
    }
}
assert.equal(k, 6, mes(4));

function findFirstGreaterThan(numbers, target) {
    for (let i = 0; i < numbers.length; i++) {
        // 如果当前元素小于或等于目标值，则跳过当前迭代
        if (numbers[i] <= target) {
            continue;
        }
        // 返回第一个大于目标值的元素
        return numbers[i];
    }
    // 如果循环完成还没有找到，返回null
    return null;
}

// 测试用例
assert.equal(findFirstGreaterThan([1, 2, 3, 4, 5], 3), 4); // 应该返回第一个大于3的值，即4
assert.equal(findFirstGreaterThan([1, 2, 3, 4, 5], 5), null); // 没有元素大于5，应返回null
assert.equal(findFirstGreaterThan([10, 20, 30], 5), 10); // 第一个元素就满足条件，应返回10
