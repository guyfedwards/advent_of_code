#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')
const inp = raw.split('\n').slice(0, -1)

const getBrackets = string => string.split('[').map(s => s.split(']'))

const extractBrackets = arr => arr.slice(1).map(v => v[0])
const extractOutBrackets = arr => arr[0].concat(arr.slice(1).map(v => v[1]))

const createSlices = s => s.split('').map((v, i) => s.substr(i, 4)).filter(v => v.length >= 4)

const predicate = s => {
    const l = s.split('')
    return s[0] === s[3] && s[1] === s[2] && s[0] !== s[1]
}

const res = inp.filter(line => {
    const outBrackets = extractOutBrackets(getBrackets(line))
    const brackets = extractBrackets(getBrackets(line))

    const outSlices = outBrackets.map(createSlices)
    const bracketSlices = brackets.map(createSlices)

    const matches = outSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicate)
    const bracketMatches = bracketSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicate)

    return matches.length > 0 && bracketMatches.length < 1
})

console.log(res.length)
