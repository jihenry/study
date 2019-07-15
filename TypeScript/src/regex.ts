// let trigger = "afdflalfjalfa.json";
// let regexp = new RegExp('^.*\.js$');
let trigger = "res/afdflalfjalfa.json";
let regexp = /^res\/.*/;
// let trigger = "res/afdflalfjalfa.json";
// let regexp = new RegExp('^res/.*[.](json|png)$');
let test = regexp.test(trigger);
console.log(test,trigger.match(regexp));
regexp = /.*[.]js$/;
test = regexp.test(trigger);
console.log(test,trigger.match(regexp));
// const s = '123-4567-89';
// const r = /\d{3}-(\d{4})-\d{2}/g;
// // console.log(r.test(s));
// const m = r.exec(s);
// if (m !== null) {
//     m.forEach((value, index) => console.log(`group ${index} : ${value}`));
// }

// const r2 = /(\d+)-(\d+)-(\d+)/g;
// console.log(s.replace(r2, '$3-$1-$2'));
 
// const r3 = /\d+/g;
// const s3 = s.replace(r3, substring => substring.split('').reverse().join(''));
// console.log(s3);
 
// const r4 = /%(?:begin|next|end)%/g;
// const s4 = '%begin%hello%next%world%end%';
// console.log(s4.split(r4).join(','));