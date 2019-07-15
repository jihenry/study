enum Color {Red = 1, Green, Blue = 4}
let c1ddd: Color = Color.Green;  // 2
let colorName: string = Color[parseInt(c1ddd.toString(),10)];  // 'Green'

console.log(c1ddd, colorName)