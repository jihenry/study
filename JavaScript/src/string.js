var a = "dgagahaha";
console.log("dffggggggggggg:", a)


function Student({
    name = "aaaaaa",
    age = 12
}) {
    this.name = name;
    console.log(name, age)

    return 111;
}

var wang = new Student({
    age: "xxxxx"
})

Student({
    name: "gggg"
})

console.log("name:", name)

console.log("student ", wang)
console.log("student ", wang)

console.log("ttxxxxx".startsWith("tt"))

function stringToBytes(str) {

    var bytes = new Array();

    var len, c;

    len = str.length;

    for (var i = 0; i < len; i++) {

        c = str.charCodeAt(i);

        if (c >= 0x010000 && c <= 0x10FFFF) {

            bytes.push(((c >> 18) & 0x07) | 0xF0);

            bytes.push(((c >> 12) & 0x3F) | 0x80);

            bytes.push(((c >> 6) & 0x3F) | 0x80);

            bytes.push((c & 0x3F) | 0x80);

        } else if (c >= 0x000800 && c <= 0x00FFFF) {

            bytes.push(((c >> 12) & 0x0F) | 0xE0);

            bytes.push(((c >> 6) & 0x3F) | 0x80);

            bytes.push((c & 0x3F) | 0x80);

        } else if (c >= 0x000080 && c <= 0x0007FF) {

            bytes.push(((c >> 6) & 0x1F) | 0xC0);

            bytes.push((c & 0x3F) | 0x80);

        } else {

            bytes.push(c & 0xFF);

        }

    }

    return bytes;

}

function byteToString(byte) {

    if (typeof byte === 'string') {

        return byte;

    }

    var str = '',

        _arr = byte;

    for (var i = 0; i < _arr.length; i++) {

        var one = _arr[i].toString(2),

            v = one.match(/^1+?(?=0)/);

        if (v && one.length == 8) {

            var bytesLength = v[0].length;

            var store = _arr[i].toString(2).slice(7 - bytesLength);

            for (var st = 1; st < bytesLength; st++) {

                store += _arr[st + i].toString(2).slice(2);

            }

            str += String.fromCharCode(parseInt(store, 2));

            i += bytesLength - 1;

        } else {

            str += String.fromCharCode(_arr[i]);

        }

    }

    return str;

}

byteArr = stringToBytes("--hcffJ7o03R5y_aAsad55VPCgSzdxUaCjz0BQ==")
console.log(byteArr)
byteStr = byteToString(byteArr)
console.log(byteStr)
var str = '怪诞咖啡';
// 字符串转base64
function encode(str) {
    // 对字符串进行编码
    var encode = encodeURI(str);
    // 对编码的字符串转化base64
    var base64 = btoa(encode);
    return base64;
}
// base64转字符串
function decode(base64) {
    // 对base64转编码
    var decode = atob(base64);
    // 编码转字符串
    var str = decodeURI(decode);
    return str;
}

// console.log(encode(str)); //JUU2JTgwJUFBJUU4JUFGJTlFJUU1JTkyJTk2JUU1JTk1JUEx
// console.log(decode(encode(str))); //怪诞咖啡


function truncated(str, num) {
    let s = '';
    for (let v of str) {
        s += v;
        num--;
        if (num <= 0) {
            break;
        }
    }
    return s;
}

function length(str) {
    let length = 0
    for (let v of str) {
        length++
    }
    return length;
}

console.log(length("王冬春abc11"))