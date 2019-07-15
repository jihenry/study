// let http = require('http');
// import http from "http"
// function test() {
//     http.get('http://www.baidu.com', function (res:{[key:string]:any}) {
//         console.log('the first Status :' + res.statusCode);
//         http.get('http://www.163.com', function (res2:{[key:string]:any}) {
//             console.log('the second Status :' + res2.statusCode);
//             http.get('http://www.sina.com', function (res3:{[key:string]:any}) {
//                 console.log('the third Status :' + res3.statusCode);
//             })
//         })
//     });
// }

// test();

import http = require('http');
import { resolve } from 'url';

class httpAsync {
    constructor() {
    }

    public async GetAsync(url: string): Promise<http.IncomingMessage> {
        var promise = new Promise<http.IncomingMessage>(resolve => {
            http.get(url, res => {
                resolve(res);
            });
        });

        return promise;
    }
}

async function test():Promise<void> {
    let ha = new httpAsync();

    let res = await ha.GetAsync("http://www.baidu.com");
    console.log('the first Status :' + res.statusCode);
    res = await ha.GetAsync("http://www.163.com");
    console.log('the first Status :' + res.statusCode);
    res = await ha.GetAsync("http://www.sina.com");
    console.log('the first Status :' + res.statusCode);
    let a = await test1();
    
    return new Promise<void>((resolve)=>{
        resolve();
    })
}

function test1():Promise<number> {
    return new Promise<number>((resolve)=>{
        resolve(2)
    })
}


