#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')

const inp = raw.split('\n')
const inp2 = inp.map(v => v.split(' ').filter(Boolean))
let col = []

for (var i = 0; i < 3; i++) {
  col.push(inp2.map(v => v[i]))
}

let all = col.reduce((acc, cur) => acc.concat(cur))
let inc = 0
let triangles = []

while (all.length > 0) {
    triangles[inc] = all.splice(0,3)
    inc++
}

const result = triangles.filter(isTriangle)

function isTriangle(t) {
    const first = parseInt(t[1]) + parseInt(t[2]) > parseInt(t[0])
    const second = parseInt(t[2]) + parseInt(t[0]) > parseInt(t[1])
    const third = parseInt(t[0]) + parseInt(t[1]) > parseInt(t[2])

    if (first && second && third) {

        return true
    }
    return false
}
