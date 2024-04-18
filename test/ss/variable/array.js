import assert from "assert";

console.log("测试语法:: 变量声明");

let a = [1, 2, 3];
assert.ok(a[0] == 1);
assert.ok(a[1] == 2);
assert.ok(a[2] == 3);
assert.ok(a[3] == undefined);
assert.ok(a.length == 3);
let b = [1, 2, 3, 3, "a", 4, 3];
let x = b.filter((value, index, arr) => {
    return value > 3;
})

assert.ok(delete a[4] == false);

assert.ok(delete a[0] == true);

assert.deepEqual(a, [undefined, 2, 3]);

const numbers = [1, 2, 3, 4];
const doubled = numbers.map((number, index, a) => {
    return number * 2;
});
assert.deepEqual(doubled, [2, 4, 6, 8]);

a = [1,2];
let c = [3,4,5]
a.push(...c)
assert.deepEqual(a, [1,2,3,4,5]);



