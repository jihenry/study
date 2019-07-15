let numberRegexp = /^[0-9]+$/;
export class ZipCodeValidator {
    isAcceptable(s: string) {
        return s.length === 5 && numberRegexp.test(s);
    }
}
// export = ZipCodeValidator;
export let AppConfig = {
    a: 111,
    b: "dddd"
}

interface ISocket {
    (a: string, b: number): void;
    sub(a: number, b: number): number;
}

export let Socket = <ISocket>function (a: string, b: number) { }
Socket.sub = function (a: number, b: number): number {
    return 1
}

const arrayd = ["a","b"]

console.log(arrayd.indexOf("v"));