const fs = require('fs')

const input = fs.readFileSync('../input.txt', 'utf8').split('\n')

function checkSlope(all, h, v) {
  let result = 0
  let horizontal = 0

  for (let rowIndex = 0; rowIndex < all.length; rowIndex += v) {
    const row = all[rowIndex + v]
    if (!row) {
      break
    }

    horizontal = horizontal + h
    const position = row[horizontal % row.length]

    if (position === '#') {
      result++
    }
  }

  return result
}

function getResultOne(all) {
  return checkSlope(all, 3, 1)
}

function getResultTwo(all) {
  const slopes = [
    [1, 1],
    [3, 1],
    [5, 1],
    [7, 1],
    [1, 2]
  ]
  let trees = []

  for (const slope of slopes) {
    trees.push(checkSlope(all, ...slope))
  }

  return trees.slice(1).reduce((total, val) => total * val, trees[0])
}

console.log('Result1: ', getResultOne(input))
console.log('Result2: ', getResultTwo(input))
