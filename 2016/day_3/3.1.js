#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')

const inp = raw.split('\n')

console.log(inp.filter(isTriangle).length)

function isTriangle(tr) {
    const t = tr.split(' ').filter(Boolean)
    const first = parseInt(t[1]) + parseInt(t[2]) > parseInt(t[0])
    const second = parseInt(t[2]) + parseInt(t[0]) > parseInt(t[1])
    const third = parseInt(t[0]) + parseInt(t[1]) > parseInt(t[2])

    if (first && second && third) {

        return true
    }
    return false
}
