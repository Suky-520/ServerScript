import assert from "assert";

console.log("测试语法::代码块");

let a = 1;

{
    a = 2;
}

assert.equal(a,2)