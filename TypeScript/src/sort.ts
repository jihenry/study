import { log } from "util";

type sortType = {enable?:number;id:number};
const sortArr:sortType[] = [
    {enable:1,id:2},
    {enable:1,id:1},
    {id:4}
];

sortArr.sort((a:sortType, b:sortType):number => {
    if (a.enable !== b.enable) {
        return a.enable ? 0: 1;
    }

    return a.id <= b.id ? 0 : 1;
});

console.log(sortArr.toString());

sortArr.sort((a:sortType, b:sortType):number => {
    if (a.enable !== b.enable) {
        return a.enable ? 0: 1;
    }

    return a.id <= b.id ? 0 : 1;
});

console.log(sortArr.toString());

type FileType = { path: string; name: string };
const files:FileType[] = [
    {path:"",name:"ddd_20190504_102010.log"},
    {path:"",name:"ddd_20190504_112310.log"},
    {path:"",name:"ddd_20190505_102012.log"},
    {path:"",name:"ddd_20190507_202010.log"}
];
const listfiles = [
    "/data/user/0/com.grhd.qpddd/files/log/./",
    "/data/user/0/com.grhd.qpddd/files/log/../",
    "/data/user/0/com.grhd.qpddd/files/log/ddd_20190704_101438.log"
]
for (const file of listfiles) {
    const fileds = file.split("/");
    if (fileds.length > 0 && (fileds[fileds.length - 1] === "." || fileds[fileds.length - 1] === "..")) {
        continue;
    }
    files.push({ name: file[fileds.length - 1], path: file });
}
files.sort((a: FileType, b: FileType) => {
    const afileds = a.name.split(/[_.]/);
    const bfileds = b.name.split(/[_.]/);
    if (afileds[1] < bfileds[1]) {
        return 0;
    } else if (afileds[1] === bfileds[1]) {
        return afileds[2] <= bfileds[2] ? 0 : 1;
    } else {
        return 1;
    }
});
console.log(files);