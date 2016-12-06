#!/usr/bin/env node
'use strict'
const md5 = require('blueimp-md5')

const raw = `uqwqemis`

let result = []
let pos = []

const checkZeros = s => /^00000/g.test(s)

const getLetter = s => s[6]

const getPos = s => s[5]

const validPos = p => p < 8 && pos.indexOf(p) == -1

for (let i = 1; i > 0; i++) {
    const hash = md5(raw + i)
    const p = getPos(hash)
    if (checkZeros(hash) && validPos(p)) {
        pos.push(p)
        result.push([p, getLetter(hash)])
        console.log(result)

        if (result.length > 7) {
            console.log(result.sort(compare).map(v => v[1]).join(''))
            break
        }
    }
}

function compare(a, b) {
    if (a[0] < b[0]) {
        return -1
    }

    if (a[0] > b[0]) {
        return 1
    }

    return 0
}
