// import zip = require("./lang");
import {AppConfig as config, ZipCodeValidator as zip} from "./lang" 

// Some samples to try
let strings = ["Hello", "98052", "101"];

// Validators to use
let validator = new zip();

console.log("ddddd")

config.b = "444444444"
console.log(config)

import moduleA from "./module" 
moduleA()

console.log(2/0)
// Show whether each string passed each validator
strings.forEach(s => {
  console.log(`"${ s }" - ${ validator.isAcceptable(s) ? "matches" : "does not match" }`);
});