// class Person {
//     public name: string;

//     constructor() {
//         this.name = "dd";
//     }
// }



// class Main {
//     private mPerson: Person;
//     constructor() {
//         this.mPerson = new(Person)
//     }
// }

function modify(p: {[key: string]:Object}) {
    console.log(p);
    let bbb = p;
    bbb.name = "nnnnnnnnnnnn";
}

let person: {[key: string]:Object} = {
    hero: {
        name: "ghhhhhh"
    }
}

console.log(person);
modify(<{[key: string]:Object}>person.hero);

console.log(person);
