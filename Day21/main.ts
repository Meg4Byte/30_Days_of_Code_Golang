'use strict';

process.stdin.resume();
process.stdin.setEncoding('utf-8');
let inputString: string = '';
let inputLines: string[] = [];
let currentLine: number = 0;
process.stdin.on('data', function(inputStdin: string): void {
    inputString += inputStdin;
});

process.stdin.on('end', function(): void {
    inputLines = inputString.split('\n');
    inputString = '';
    main();
});

function readLine(): string {
    return inputLines[currentLine++];
}

function main() {
    var numberList: number[] = new Array(Number(inputLines[0]));
    for (let i = 1; i < Number(inputLines[0]) +1 ;i++ ){
        numberList.push(Number(inputLines[i]))
    }
    var stringList: string[] = new Array(Number(inputLines[Number(inputLines[0]) +1])); ;
    for (let i = Number(inputLines[0]) +2; i < Number(inputLines.length) ;i++ ){
        stringList.push(inputLines[i])
    }
    printArray(numberList)
    printArray(stringList)
}
function printArray(array: any[]) {
    array.forEach(e => {
        console.log(e)
    });
}
