import assert from "assert"

console.log("测试语法:: try catch");

try{
    throw "abc";
    assert.ok(false);
}catch (e){
    assert.equal(e, "abc");
}finally {
    assert.ok(true);
}

try{
    throw new Error("abc");
    assert.ok(false);
}catch (e){
    assert.equal(e.name, "Error");
}finally {
    assert.ok(true);
}

try{
    throw new TypeError("abc");
    assert.ok(false);
}catch (e){
    assert.equal(e.name, "TypeError");
}finally {
    assert.ok(true);
}

try{
    assert.ok(true);
}catch (e){
    assert.ok(false);
}finally {
    assert.ok(true);
}