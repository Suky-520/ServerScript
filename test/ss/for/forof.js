import assert from "assert"

console.log("测试语法 forof");

let obj = {
    name: "Xx",
    age: 12
}
let objKey = ['name', 'age'];
let objValue = ['Xx', 12];
let arr = [2, 3, 4];

let i = 0;
for (let x of arr) {
    assert.equal(x, arr[i]);
    i++;
}

i = 0;
for (let [k, v] of arr) {
    assert.equal(k, i);
    assert.equal(v, arr[i]);
    i++;
}

i = 0;
for (let x of obj) {
    assert.equal(x, objValue[i]);
    i++;
}

i = 0;
for (let [k, v, x] of obj) {
    assert.equal(k, objKey[i]);
    assert.equal(v, objValue[i]);
    assert.equal(x, undefined);
    i++;
}

let str = "abc123";
let strArr = ["a", "b", "c", "1", "2", "3"];

i = 0;
for (let x of str) {
    assert.equal(x, strArr[i]);
    i++;
}

i = 0;
for (let [k, v, x] of str) {
    assert.equal(k, i);
    assert.equal(v, strArr[i]);
    assert.equal(x, undefined);
    i++;
}

let set = new Set();
let setArr = ["1", "2", "3"];
set.add("1");
set.add("2");
set.add("3");

i = 0;
for (let [k, v, x] of set) {
    assert.equal(k, setArr[i]);
    assert.equal(v, undefined);
    assert.equal(x, undefined);
    i++;
}

i = 0;
for (let x of set) {
    assert.equal(x, setArr[i]);
    i++;
}

let map = new Map();
let mapKeyArr = ["1", "2", "3"];
let mapValArr = ["a", "b", "c"];
map.set("1", "a");
map.set("2", "b");
map.set("3", "c");

i = 0;
for (let [k, v, x] of map) {
    assert.equal(k, mapKeyArr[i]);
    assert.equal(v, mapValArr[i]);
    assert.equal(x, undefined);
    i++;
}

i = 0;
for (let x of map) {
    assert.equal(x, mapValArr[i]);
    i++;
}