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
    // len of Array: inputLines[0]
    // string like '1 2 3': inputLines[1] 
    
    
    const arrayLen = Number(inputLines[0])
    
    const numberArray: number[]= [];
    inputLines[1].split(' ').forEach( x => { numberArray.push(Number(x))});
    let maxNum: number;
    numberArray.sort((n1,n2) => n1 - n2);
    maxNum = numberArray[numberArray.length -1] - numberArray[0]
    if (maxNum < 0) {
        maxNum = -maxNum
    } 
    
    console.log(maxNum)
}
