'use strict';
import ByteBuffer = require("./bytebuffer");
var fs = require('fs');

const data = new ByteBuffer(10);

// console.log(data.mark(6));
data.append("abcdefgrtuuuuu");
console.log(data.toArrayBuffer(),data.toString());
console.log(new Uint8Array(data.view))
const data1 = data.copy(0,data.offset);
console.log(data.offset,data.toArrayBuffer().byteLength);
let addd = "11111111111111";
if (addd.length > data.limit) {
   addd = addd.substr(addd.length-data.limit,data.limit); 
}
const leftLength = data.limit - data.offset;
if (addd.length > leftLength) {
    const surplus = addd.length - leftLength;
    console.log(data.toArrayBuffer(),data.toString());
    if (surplus >= data.limit) {
        data.clear();
    } else {
        data.compact(surplus)
        data.resize(10); //需要修改容量
        data.limit = 10; //需要修改limit
    }
    data.append(addd)
}
console.log(data.toArrayBuffer()); //toArrayBuffer表示剩余多少字节
fs.writeFileSync("D:/1.txt",data.view)
console.log(data.view)
fs.writeFileSync("D:/2.txt",data.view.subarray(0,data.offset))

// const mMaxLogSize = 10;
// const mBuffer = new ByteBuffer(mMaxLogSize);
// mBuffer.append("点点点")
// let content = `附件六大解放辣椒干\n`;
// const byteLength = ByteBuffer.fromUTF8(content).view.byteLength;
// if (byteLength > mMaxLogSize) {
//     content = content.substr(content.length-mMaxLogSize,mMaxLogSize);
// }

// const leftLength = mBuffer.limit - mBuffer.offset;
// if (content.view.length > leftLength) {
//     const surplus = content.view.length - leftLength;
//     if (surplus >= mMaxLogSize) {
//         mBuffer.clear();
//     } else {
//         mBuffer.compact(surplus, mMaxLogSize);
//         mBuffer.resize(mMaxLogSize); //需要修改容量
//         mBuffer.limit = mMaxLogSize; //需要修改limit，只是compact
//     }
// }
// mBuffer.append(content);
// console.log(mBuffer.view.subarray(0,mBuffer.offset));
// console.log(mBuffer);