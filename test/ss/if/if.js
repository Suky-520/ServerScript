import assert from "assert"

console.log("测试语法::if");

function mes(i) {
    return `测试语法::if  测试用例[${i}]`
}

let a = 1;
if (a == 1) {
    assert.ok(true, mes(1));
} else {
    assert.ok(false, mes(2));
}

if (a == 2) {
    assert.ok(false, mes(3));
} else if (a == 1) {
    assert.ok(true, mes(4));
} else {
    assert.ok(false, mes(5));
}

if (a == 2) {
    assert.ok(false, mes(6));
} else if (a == 3) {
    assert.ok(false, mes(7));
} else {
    assert.ok(true, mes(8));
}