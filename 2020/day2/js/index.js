const fs = require('fs')

const input = fs.readFileSync('../input.txt', 'utf8').split('\n')

function parseEntry(raw) {
  const parts = raw.split(/[-\s:]/).filter(Boolean);
  if (parts.length === 0) {
    return null
  }

  return {
    min: parts[0],
    max: parts[1],
    letter: parts[2],
    password: parts[3]
  }
}

function getResultOne(all) {
  let result = 0;

  for (const e of all) {
    const entry = parseEntry(e)
    if (!entry) {
      continue
    }

    const count = entry.password.split('').filter(v => v === entry.letter).length

    if (count >= entry.min && count <= entry.max) {
      result++
    }
  }

  return result
}

function getResultTwo(all) {
  let result = 0;

  for (const e of all) {
    const entry = parseEntry(e)
    if (!entry) {
      continue
    }

    const max = entry.password[entry.max-1]
    const min = entry.password[entry.min-1]


    if ((max === entry.letter && min !== entry.letter) || (max !== entry.letter && min === entry.letter)) {
      result++
    }
  }

  return result
}

console.log('Result1: ', getResultOne(input))
console.log('Result2: ', getResultTwo(input))
