
// map filter reduce array

const test = [1, 2, 3, 4, 5, 6, 7, 8, 9, 10];

const result = test.map((item) => item * 2).filter((item) => item > 10).reduce((acc, item) => acc + item, 0);
console.log(result);
