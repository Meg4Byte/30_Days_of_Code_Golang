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

export interface MyBookProps {
     title : string;
     author : string;
     price : number;
}

export abstract class BaseBook {
    readonly title : string;
    readonly author : string;
    readonly price : number;
    constructor(props: MyBookProps) {
        this.title = props.title
        this.author = props.author
        this.price = props.price
    }
    public display() :void {
        console.log(`Title: ${this.title}`);
        console.log(`Author: ${this.author}`);
        console.log(`Price: ${this.price}`);
    }
}


export class MyBook extends BaseBook {
    constructor(props: MyBookProps) {
        super(props);
    }
}

function main() {
    const mybook = new MyBook({title: inputLines[0], author: inputLines[1], price: Number(inputLines[2])})
    mybook.display()
}
