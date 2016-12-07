#!/usr/bin/env node
'use strict'
const fs = require('fs')
const { countLetters } = require('../utils')

const raw = fs.readFileSync(process.argv[2], 'utf8')
const inp = raw.split('\n')

// for (let i = 0; i < inp.length; i++) {
//     inp.map()
// }
const createLetterArray = s => s.split('').map(l => l)

const res = inp.map(createLetterArray)

// console.log(res)

const createCols = input => input.reduce((acc, cur) => {

    createLetterArray(cur).map((l, i) => {
        if (typeof acc[i] != 'object') {
            acc[i] = []
        }

        acc[i].push(l)
    })

    return acc
}, [])

const letterCounts = createCols(inp).map(i => countLetters(i.join('')))

const getHighest = counts => counts.reduce((acc, cur, i) => {
    Object.keys(cur).map(key => {
        if (acc[i] === undefined || cur[key] > acc[i][1]) {
            acc[i] = [key, cur[key]]
        }
    })

    return acc
}, [])

const getLowest = counts => counts.reduce((acc, cur, i) => {
    Object.keys(cur).map(key => {
        if (acc[i] === undefined || cur[key] < acc[i][1]) {
            acc[i] = [key, cur[key]]
        }
    })

    return acc
}, [])

const flatten = arr => arr.map(v => v[0]).join('')

console.log(flatten(getHighest(letterCounts)))
console.log(flatten(getLowest(letterCounts)))



