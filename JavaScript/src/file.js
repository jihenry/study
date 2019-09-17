var fs = require("fs")

fs.writeFileSync("D:\\1.txt","xxxxxxxxxxxxxxxxxxxxxxx")

var data = fs.readFileSync("D:\\workspace\\minigame\\Client\\assets\\scripts\\ui\\subpackage.meta","utf8")
let json = JSON.parse(data);
console.log("json:",json)