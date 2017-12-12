#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')
const inp = raw.split('\n').slice(0, -1)

const getBrackets = string => string.split('[').map(s => s.split(']'))

const extractBrackets = arr => arr.slice(1).map(v => v[0])
const extractOutBrackets = arr => arr[0].concat(arr.slice(1).map(v => v[1]))

const createSlices = count => s => s.split('').map((v, i) => s.substr(i, count)).filter(v => v.length >= count)

const predicateAbba = s => {
    const l = s.split('')
    return s[0] === s[3] && s[1] === s[2] && s[0] !== s[1]
}

const predicateAba = s => {
    const l = s.split('')
    return s[0] === s[2] && s[0] !== s[1]
}

const bab = s => s[1] + s[0] + s[1]

const res1 = inp.filter(line => {
    const outBrackets = extractOutBrackets(getBrackets(line))
    const brackets = extractBrackets(getBrackets(line))

    const outSlices = outBrackets.map(createSlices(4))
    const bracketSlices = brackets.map(createSlices(4))

    const matches = outSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicateAbba)
    const bracketMatches = bracketSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicateAbba)

    return matches.length > 0 && bracketMatches.length < 1
})

const res2 = inp.filter(line => {
    const outBrackets = extractOutBrackets(getBrackets(line))
    const brackets = extractBrackets(getBrackets(line))

    const outSlices = outBrackets.map(createSlices(3))
    const bracketSlices = brackets.map(createSlices(3))

    const matches = outSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicateAba)
    const bracketMatches = bracketSlices.reduce((acc, cur) => acc.concat(cur), []).filter(predicateAba)
    const bracketReverse = bracketMatches.map(bab)

    const r = matches.filter(v => bracketMatches.indexOf(bab(v)) != -1)

    return r.length > 0
})


console.log(res1.length)
console.log(res2.length)
