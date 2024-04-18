import assert from "assert"
let a ="hello"

export default {
    name:"zhangsan",
    getName1(){
        let a="lisi"
        return a+this.name
    },
    getName2(){
        return a
    },
    getName3(){
        let a="lisi"
        return a
    },
}