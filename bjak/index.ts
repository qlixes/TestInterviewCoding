// from gorilla test

function sortTypes(arr: any[]) {

    let result: any[] = [];
    let alphas: string[] = [];
    let numbers: number[] = [];
    let defaults: any[] = [];
    let objects: object[] = [];

    for (let val of arr) {

        let dataType = typeof val;

        switch (dataType) {

            case 'object':
                for (let obs of val) {
                    if (objects.indexOf(obs) === -1) {
                        objects.push(obs);
                    }
                }
                break;

            case 'number':
                if (numbers.indexOf(val) === -1) {
                    numbers.push(val);
                }
                break;

            case 'string':
                if (alphas.indexOf(val) === -1) {
                    alphas.push(val);
                }
                break;

            default:
                if (defaults.indexOf(val) === -1) {
                    defaults.push(val);
                }
        }
    }

    if (numbers.length > 0) {
        result.push(numbers);
    }

    if (alphas.length > 0) {
        result.push(alphas);
    }

    if (objects.length > 0) {
        result.push(objects);
    }

    if (defaults.length > 0) {
        result.push(defaults);
    }

    return JSON.stringify(result);
}

// unit test
let test1 = [1, 2, "hello"];
console.log(sortTypes(test1));
let test2 = [1, 1, [1, 2, 3]];
console.log(sortTypes(test2));
// let test3 = [null, null];
// console.log(sortTypes(test3));
let test4 = [[1, 2, 3]];
console.log(sortTypes(test4));
