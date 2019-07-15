import { AppConfig as config } from "./lang"

const sym1 = Symbol()
let c = "1111"
let a = () => {
    console.log("b=" + config.b)
    console.log("b=" + config.b)
}
let bc = 1
console.log("bc=", bc)

let result1 = {
    sym1: "dddgggg"
}
console.log("result1=", result1.sym1)

interface Map {
    [key: string]: any
}

class C {
    private dd: string[] = ["dd","666","777"]
    constructor() {
    }
    ["ccc"]() {
        return "lllllllllllllll"
    }
    [Symbol.iterator]() {
        const self = this;
        let index = 0;
        return {
            next() {
                if (index < self.dd.length) {
                    return {
                        value: self.dd[index++],
                        done: false
                    }
                } else {
                    return { value: undefined, done: true };
                }
            }
        };
    }
}

let c1 = new C()
for (const cins of c1) {
    console.log("iter:" + cins)
}
console.log("c['ccc']" + c1["ccc"]())

let g = <Map>{}
g["ddd"] = 1

for (const key in g) {
    console.log(g[key])
}

let result = ["dd", "111", "ccc"]
for (const pet of result) {
    console.log(pet)
}

// let pets = new Set(["Cat","Dog","ddd"])
// pets.add("Dog")
// for (const pet of pets) {
//     console.log(pet)
// }

export = a