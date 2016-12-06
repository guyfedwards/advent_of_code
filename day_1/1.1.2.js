#!/usr/bin/env node
'use strict';
const fs = require('fs')

const input = fs.readFileSync(process.argv[2], 'utf8').split(', ')

var direction = 'up'
let vertical = 0
let horizontal = 0
let points = []

input.forEach(function(i) {
    const dir = i[0]
    const dis = parseInt(i.slice(1))

    switch(direction) {
        case 'up':
            for (var i = 0; i < dis; i++) {
                if (dir == 'R'){
                    horizontal += 1
                    direction = 'right'
                } else {
                    horizontal -= 1
                    direction = 'left'
                }
                addPoints(horizontal, vertical)
            }
            break

        case 'right':
            for (var i = 0; i < dis; i++) {
                if (dir == 'R') {
                    vertical -= 1
                    direction = 'down'
                } else {
                    vertical += 1
                    direction = 'up'
                }
                addPoints(horizontal, vertical)
            }
            break

        case 'down':
            for (var i = 0; i < dis; i++) {
                if (dir == 'R') {
                    horizontal -= 1
                    direction = 'left'
                } else {
                    horizontal += 1
                    direction = 'right'
                }
                addPoints(horizontal, vertical)
            }
            break

        case 'left':
            for (var i = 0; i < dis; i++) {
                if (dir == 'R') {
                    vertical += 1
                    direction = 'up'
                } else {
                    vertical -= 1
                    direction = 'down'
                }
                addPoints(horizontal, vertical)
            }
            break

        default:
            break
    }
})


const result2 = points.map(c => points.filter(x => c[0] === x[0] && c[1] === x[1])).filter(x => x.length === 2)

function addPoints(x, y) {
    points.push([x, y])
}

console.log(`1::::: H: ${horizontal}`, `V: ${vertical}`, `D: ${Math.abs(vertical) + Math.abs(horizontal)}`)
console.log('2:::::', Math.abs(result2[0][0][0]) + Math.abs(result2[0][0][1]))

