function Student(name) {
    this.name = name;
    this.hello = function () {
        console.log('Hello, ' + this.name + '!');
    }
}
// var testobj = {

//     name: 1111,
//     meth: function() {
//         console.log("ddddgagag")
//     },
//     prototype: {
//         constructor: Student,
//         tetproto: function() {
//             console.log("ddggggggggggggggg")
//         }
//     }
// }
// var testobj1 = new testobj();

var arr = [1,2,3]

class ddd {
    constructor() {
        this.mea = "1";
    }
    hello(){
        console.log("dddd");
    }
}

class A extends ddd {
    constructor(name, grade) {
        super(name); // 记得用super调用父类的构造方法!
        this.grade = grade;
    }

    myGrade() {
        alert('I am at grade ' + this.grade);
    }
}

var AObj = new A();
var xiaoming = new Student('小明');
var obj1 = new Object();
var obj2 = new Object();
console.log(obj1.name);
xiaoming.name; // '小明'
xiaoming.hello(); // Hello, 小明!

console.log(xiaoming);