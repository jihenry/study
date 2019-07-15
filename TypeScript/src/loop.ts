// let hh1 = {a:1,b:"12"};
// let hh2 = {a:2,b:"52"};
// let hh3 = {a:3,b:"42"};
let h = ["2","4","5"];
let h1 = h.slice();
console.log(h1);
for (const hn of h) {
    console.log(`hn:${hn}`)
    const index = h1.indexOf(hn);
    h1.splice(index,1);
    console.log(`h1:${h1}`)
    console.log(`h:${h}`)
}