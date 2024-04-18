import assert from "assert"

let a ="hello"

function b(){
    let a = "vino"
    return {
        getName(){
            let a="vvv"
            return {
                getAge(){
                    return a
                }
            }
        },
        getAge(){
            return a
        }
    }
}
export default b