import assert from "assert";

console.log("测试对象:: set");

let set1 = new Set();
set1.add("key1");
set1.add("key2");
set1.add("key3");
set1.add("key4");
assert.ok(set1.size == 4);
set1.delete("key")
assert.ok(set1.size == 4);
assert.ok(set1.has("key3"));
set1.delete("key1")
assert.ok(set1.size == 3);
set1.clear()
assert.ok(set1.size == 0);
