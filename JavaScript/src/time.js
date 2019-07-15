function printTime() {
    console.log('It is time!');
}

setTimeout(printTime, 1000);
console.log(this)
console.log('done');