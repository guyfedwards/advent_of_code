import Data.Char
import Data.List
import Debug.Trace

toInt :: String -> Int
toInt str = read str::Int

clean :: String -> String
clean str = concatMap (\c -> if c == '+' then []; else [c]) str

numTimeFound :: Int -> [Int] -> Int
numTimeFound x list = length $ filter (== x) list

moreThanOnce :: [Int] -> (Int, Int) -> Bool
moreThanOnce l (x, i) = do
  let slice = take (i + 1) l
  let numfound = numTimeFound x slice
  numfound > 1

main = do
  contents <- readFile "./input.txt"
  let ints = map toInt . map clean $ lines contents

  let sum = foldl (+) 0 ints
  let freq = scanl (+) 0 (take 200000 . cycle $ ints)
  let indFreq = zip freq [0..]

  let a = take 10 . cycle $ [1, 2, 3, 4, 5]
  let inda = zip [0..] a

  print (head (filter (moreThanOnce freq) indFreq))


--   ANSWER: 245
