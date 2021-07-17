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
 
interface CalculatorProps {
    readonly numberN :number;
    readonly numberP :number;
}
class Calculator {
    isNegative: boolean = false;
    constructor(props: CalculatorProps ) {
        this.checkNotNegative(props.numberN, props.numberP);
        if (!this.isNegative) {
            console.log(props.numberN ** props.numberP)
        }
    }
    private checkNotNegative(n: number , p: number){
        if ( n < 0 || p < 0) {
            console.log('n and p should be non-negative')
            this.isNegative = true;
        }
    }
}

function main() {
    for (let i = 1 ; i < inputLines.length; i++) {
        let numbers = inputLines[i].split(' ');
        new Calculator({numberN: Number(numbers[0]),numberP: Number(numbers[1]) })
    }
}