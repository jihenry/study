var Base64 = require('js-base64').Base64;


console.log(Base64.encode('dankogai'));  // ZGFua29nYWk=
console.log(Base64.encode('小飼弾'));    // 5bCP6aO85by+
console.log(Base64.encodeURI('小飼弾')) // 5bCP6aO85by-
console.log(Base64.decode("eyJtYXhpbmRleCI6NCwiY2hpbGRzIjpbeyJpZCI6IjEiLCJuaWNrIjoiMSIsInByb3RyYWl0IjoiIiwicGhvbmVubyI6IjEzMTI0NjMwMzgxIiwidGltZSI6MSwibGV2ZWwiOjAsImlzcmVhbCI6dHJ1ZSwid3giOiIxIiwicXEiOiIxIn1dfQ=="))
console.log(atob("eyJtYXhpbmRleCI6NCwiY2hpbGRzIjpbeyJpZCI6IjEiLCJuaWNrIjoiMSIsInByb3RyYWl0IjoiIiwicGhvbmVubyI6IjEzMTI0NjMwMzgxIiwidGltZSI6MSwibGV2ZWwiOjAsImlzcmVhbCI6dHJ1ZSwid3giOiIxIiwicXEiOiIxIn1dfQ=="))