#!/usr/bin/env node
'use strict'
const fs = require('fs')

const raw = fs.readFileSync(process.argv[2], 'utf8')

const ins = raw.split('\n')

const numbers = [['', '', 1, '', ''], ['', 2, 3, 4, ''], [5, 6, 7, 8, 9], ['', 'A', 'B', 'C', ''], ['', '', 'D', '', '']]

let col = 0
let row = 2

function getKey(r=row, c=col) {
    return numbers[r][c]
}

function move(dir) {
    switch(dir) {
        case 'U':
            if (row != 0 && getKey(row -1, col) != '') {
                row -= 1
            }
            break
        case 'D':
            if (row != 4 && getKey(row + 1, col) != '') {
                row += 1
            }
            break
        case 'L':
            if (col != 0 && getKey(row, col - 1) != '') {
                col -= 1
            }
            break
        case 'R':
            if (col != 4 && getKey(row, col + 1) != '') {
                col += 1
            }
            break
    }
}

const result = ins.map(line => {
    line.split('').map(i => move(i))
    return getKey()
})

console.log(result)

