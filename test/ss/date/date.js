import assert from "assert";

console.log("测试对象:: date");

let d0 = new Date('2024-04-05 06:07:08', 'yyyy-MM-dd HH:mm:ss');
let d1 = new Date('2024-01-21', 'yyyy-MM-dd');
let d2 = new Date('21-21', 'HH-mm');
let d3 = new Date(1712823003);
let d4 = new Date(1712823003092);

assert.deepEqual(d0, "2024-04-05 06:07:08");
assert.deepEqual(d1, "2024-01-21 00:00:00");
assert.deepEqual(d2, "0000-01-01 21:21:00");
assert.deepEqual(d3, "2024-04-11 16:10:03");
assert.deepEqual(d4, "2024-04-11 16:10:03");
assert.deepEqual(d4.format("yyyy-MM-dd"), "2024-04-11");

assert.ok(d3.getTime() == 1712823003000);
assert.ok(d4.getTime() == 1712823003092);
assert.ok(d4.getMilliseconds() == 92);

assert.ok(d0.getFullYear() == 2024);
assert.ok(d0.getMonth() == 4);
assert.ok(d0.getDate() == 5);
assert.ok(d0.getDay() == 5);
assert.ok(d0.getHours() == 6);
assert.ok(d0.getMinutes() == 7);
assert.ok(d0.getSeconds() == 8);
assert.ok(d0.getMilliseconds() == 0);

