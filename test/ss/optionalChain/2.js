import assert from "assert";

console.log("测试语法:: ?.");

let a = null;
assert.equal(a?.b?.c?.v(a?.n, a?.n), undefined);

a = {b: null};
assert.equal(a?.b?.c?.v(a?.n, a?.n), undefined);

a = {b: {c: null}};
assert.equal(a?.b?.c?.v(a?.n, a?.n), undefined);

a = {b: {c: {v: null}}};
assert.equal(a?.b?.c?.v?.(a?.n, a?.n), undefined);

a = {b: {c: {v: (x, y) => x + y}}};
assert.equal(a?.b?.c?.v?.(a?.n, a?.n), NaN);

a = {b: {c: {v: (x, y) => x + y}}, n: 2};
assert.equal(a?.b?.c?.v?.(a?.n, a?.n), 4);

a = {b: {c: {v: (x, y) => x + y}}, n: null};
assert.equal(a?.b?.c?.v?.(a?.n, a?.n), 0);

a = {b: {c: {v: (x, y) => x + y}}, n: 2, m: 3};
assert.equal(a?.b?.c?.v?.(a?.n, a?.m), 5);

