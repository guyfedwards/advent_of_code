const fs = require('fs')

const input = fs.readFileSync('../input.txt', 'utf8').split('\n').map(v => parseInt(v, 10));

function getResultOne(all) {
  let result;

  for (const n of all) {
    const remainder = 2020 - n
    if (all.includes(remainder)) {
      result = n * remainder
      break
    }
  }
  return result
}

function getResultTwo(all) {
  let result;

  outer:
  for (const n of all) {
    for (const nn of all) {
      for (const nnn of all) {
        if (n + nn + nnn === 2020) {
          result = n * nn * nnn
          break outer
        }
      }
    }
  }

  return result
}

console.log('Result1: ', getResultOne(input))
console.log('Result2: ', getResultTwo(input))
