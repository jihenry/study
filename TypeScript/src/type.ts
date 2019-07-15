namespace dd {
    let a:Object = true;
    let b = a.toString();
    let f:{[key:string]:string} = {
        a: "1",
        b:"d"
    }
    console.log("f="+f);
    console.log(`${b},type=${typeof b}`);
    

    let str:Object = "dddd";
    console.log(typeof str);


    function checkString(p:Object):boolean {
        return typeof p === "string" && p === "true";
    }

    console.log(checkString("false"));
    
    export const outRepeatArray = <T>(arr: T[]): T[] => {
        const ret: T[] = [];
        for (const element of arr) {
            if (ret.indexOf(element) < 0) {
                ret.push(element);
            }
        }

        return ret;
    };

    const aaaaa = [1,22,43,4,5,7]
    aaaaa.splice(-1,1)
    console.log(aaaaa);
    console.log(outRepeatArray(["dddd","1","2","2"]))
    // console.log(aaaaa.filter(x => x !== 4));
}
