
import assert from "assert"
let a = "hello"

let c ={
    name:"zhangsan",
    getName1(){
        let a="lisi"
        return a+this.name
    },
    getName2(){
        return a+"test"
    },
    getName3(){
        let a="lisi"
        return a+"test"
    },
}

assert.ok(c.getName1()=="lisizhangsan")
assert.ok(c.getName2()=="hellotest")
assert.ok(c.getName3()=="lisitest")