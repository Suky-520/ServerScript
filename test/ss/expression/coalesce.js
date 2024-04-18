import assert from "assert";

console.log("测试空值合并运算符");

let a = null;
assert.equal(a ?? 'default', 'default');

let b = undefined;
assert.equal(b ?? 'default', 'default');

let c = 0;
assert.equal(c ?? 'default', 0);

let d = '';
assert.equal(d ?? 'default', '');

let e = false;
assert.equal(e ?? 'default', false);

let f;
assert.equal(f ?? 'default', 'default');

let g = 'value';
assert.equal(g ?? 'default', 'value');

let h = NaN;
assert.equal(h ?? 'default', NaN);

let aa = null;
aa ??= 'default';
assert.equal(aa, 'default');

let bb = undefined;
bb ??= 'default';
assert.equal(bb, 'default');

let cc = 0;
cc ??= 'default';
assert.equal(cc, 0);

let dd = '';
dd ??= 'default';
assert.equal(dd, '');

let ee = false;
ee ??= 'default';
assert.equal(ee, false);

let ff;
ff ??= 'default';
assert.equal(ff, 'default');

let gg = 'value';
gg ??= 'default';
assert.equal(gg, 'value');

let hh = NaN;
hh ??= 'default';
assert.equal(hh, NaN);
