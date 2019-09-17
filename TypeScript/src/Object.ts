const func111 = (key:string,value:Object):void =>{
    console.log(key,typeof value);
}

func111("xxxx",123);
console.log(123..toString())
const cache: Object[] = [];
const keyCache: string[] = [];
let astrifify = JSON.stringify({},(key: string, value: Object):Object=>{
    if (typeof value === 'object' && value !== null) {
        if (Array.isArray(value) && value.length <= 0) {
            return value;
        }
        const index = cache.indexOf(value);
        if (index !== -1) {
            return `[Circular ${keyCache[index]}`;
        }
        cache.push(value);
        keyCache.push((key === "" || key === undefined) ? 'root' : key);
    }

    return value;
});

console.log("=========================",astrifify);
