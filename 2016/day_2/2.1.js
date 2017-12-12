#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')

const ins = raw.split('\n')

const numbers = [[1, 2, 3], [4, 5, 6], [7, 8, 9]]

let row = 1
let col = 1

function getKey(r, c) {
    return numbers[r][c]
}

function move(dir) {
    switch(dir) {
        case 'U':
            if (row != 0) {
                row -= 1
            }
            break
        case 'D':
            if (row != 2) {
                row += 1
            }
            break
        case 'L':
            if (col != 0) {
                col -= 1
            }
            break
        case 'R':
            if (col != 2) {
                col += 1
            }
            break
    }
}

const result = ins.map(line => {
    line.split('').map(i => move(i))
    return getKey(row, col)
})

console.log(result)


