#!/usr/bin/env node
'use strict';
const input = 'R1, R1, R3, R1, R1, L2, R5, L2, R5, R1, R4, L2, R3, L3, R4, L5, R4, R4, R1, L5, L4, R5, R3, L1, R4, R3, L2, L1, R3, L4, R3, L2, R5, R190, R3, R5, L5, L1, R54, L3, L4, L1, R4, R1, R3, L1, L1, R2, L2, R2, R5, L3, R4, R76, L3, R4, R191, R5, R5, L5, L4, L5, L3, R1, R3, R2, L2, L2, L4, L5, L4, R5, R4, R4, R2, R3, R4, L3, L2, R5, R3, L2, L1, R2, L3, R2, L1, L1, R1, L3, R5, L5, L1, L2, R5, R3, L3, R3, R5, R2, R5, R5, L5, L5, R2, L3, L5, L2, L1, R2, R2, L2, R2, L3, L2, R3, L5, R4, L4, L5, R3, L4, R1, R3, R2, R4, L2, L3, R2, L5, R5, R4, L2, R4, L1, L3, L1, L3, R1, R2, R1, L5, R5, R3, L3, L3, L2, R4, R2, L5, L1, L1, L5, L4, L1, L1, R1'.split(', ')

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

console.log(`1::::: H: ${horizontal}`, `V: ${vertical}`, `D: ${Math.abs(vertical) + Math.abs(horizontal)}`)

const result2 = points.map(c => points.filter(x => c[0] === x[0] && c[1] === x[1])).filter(x => x.length === 2)
console.log('2:::::', Math.abs(result2[0][0][0]) + Math.abs(result2[0][0][1]))

function addPoints(x, y) {
    points.push([x, y])
}
