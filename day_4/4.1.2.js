#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')

const inp = raw.split('\n')

const getChecksum = s => s.slice(-7).replace(/(\[|\])/g, '')

const getId = s => s.slice(-10, -7)

const getName = s => s.slice(0, -11)

const stripDashes = s => s.replace(/-/g, '')

const createTuples = result => Object.keys(result).map(l => [l, result[l]])

const sortTuples = tuples => tuples.slice().sort(compare).reverse()

const getCorrectChecksum = arr => arr.slice(0, 5).map(v => v[0]).join('')

function countLetters(string) {
    const result = {}

    string.split('').map(l => {
        if (result[l] === undefined) {
            result[l] = 1
        } else {
            result[l] += 1
        }
    })

    return result
}

function compare(a, b) {
    if (a[1] < b[1]) {
        return -1
    }
    if (a[1] > b[1]) {
        return 1
    }
    return 0
}

function resort2(arr) {
    console.log(arr)
    const max = arr[0][1]
    let result = []

    for (let i = max; i > 0; i--) {
        const cur = arr.filter(v => v[1] === i).sort()
        result.push(...cur)
    }

    return result
}

const ans = inp.filter(i => {
    const hi = getCorrectChecksum(resort2(sortTuples(createTuples(countLetters(stripDashes(getName(i)))))))
    const check = getChecksum(i)

    return check === hi
})

const sumOfIds = ans.reduce((acc, cur) => acc + parseInt(getId(cur)), 0)

const shiftLetter = count => letter => {
    const shift = count % 26
    const char = letter.charCodeAt(0)
    const newPos = (char - 96) + shift
    const newChar = newPos > 26
        ? (newPos + 96) - 26
        : newPos + 96

    return String.fromCharCode(newChar).replace(/2/g, ' ')
}

const findNorth = s => /north/.test(s)

const ans2 = inp.map(i => getName(i).split('').map(shiftLetter(getId(i))).concat(` - ${getId(i)}`).join('')).filter(findNorth)

console.log(sumOfIds)
console.log(ans2)


