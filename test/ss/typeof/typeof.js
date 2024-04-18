import assert from "assert";

console.log("测试对象:: typeof");

assert.deepEqual(typeOf("1"), {type: 1, typeString: "string", name: "string", native: true});

assert.deepEqual(typeOf(1), {type: 2, typeString: "number", name: "number", native: true})

assert.deepEqual(typeOf(true), {type: 3, typeString: "boolean", name: "boolean", native: true})

assert.deepEqual(typeOf(null), {type: 4, typeString: "null", name: "null", native: true})

assert.deepEqual(typeOf(undefined), {type: 5, typeString: "undefined", name: "undefined", native: true})

assert.deepEqual(typeOf(1n), {type: 6, typeString: "bigint", name: "bigint", native: true})

assert.deepEqual(typeOf([1, 2]), {type: 8, typeString: "array", name: "array", native: true})

assert.deepEqual(typeOf(eval), {type: 9, typeString: "function", name: "eval", native: true})

assert.deepEqual(typeOf({}), {type: 7, typeString: "object", name: "object", native: false})

assert.deepEqual(typeOf(() => {}), {type: 9, typeString: "function", name: "", native: false})

assert.deepEqual(typeOf(Date), {type: 11, typeString: "class", name: "Date", native: true})

assert.deepEqual(typeOf(new Date()), {type: 7, typeString: "object", name: "date", native: true})

assert.deepEqual(typeOf(Buffer), {type: 11, typeString: "class", name: "Buffer", native: true})

assert.deepEqual(typeOf(new Buffer()), {type: 10, typeString: "buffer", name: "buffer", native: true})