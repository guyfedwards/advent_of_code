const stripDashes = s => s.replace(/-/g, '')

function countLetters(string) {
    const result = {}

    string.split('').map(l => {
        if (result[l] === undefined) {
            result[l] = 1
        } else {
            result[l] += 1
        }
    })

    return result
}

module.exports = {
    stripDashes,
    countLetters
}
