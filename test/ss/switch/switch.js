import assert from "assert";

console.log("测试语法:: switch");

switch (1) {
    case 1:
        assert.ok(true);
    case 2:
        assert.ok(false);
    default:
        assert.ok(false);
}

switch (2) {
    case 1:
        assert.ok(false);
    case 2:
        assert.ok(true);
    default:
        assert.ok(false);
}
