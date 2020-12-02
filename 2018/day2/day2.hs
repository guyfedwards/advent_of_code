import Data.List
import Debug.Trace

clean :: String -> String
clean str = concatMap (\c -> if c == '+' then []; else [c]) str

numTimeFound :: [Char] -> Char -> Int
numTimeFound list x  = length $ filter (== x) list

reduceLine :: [Char] -> (Int,Int) -> Char -> (Int,Int)
reduceLine line (x,y) char
  | (numTimeFound line char) > 2 = (x, y + 1)
  | (numTimeFound line char) > 1 = (x + 1, y)
  | otherwise = (x,y)

parseLine :: [(Int,Int)] -> String -> [(Int,Int)]
parseLine acc line = do
  let uniqs = nub line

  let total = foldl (reduceLine line) (0,0) uniqs

  -- let moreThanOnce = [[(x,y) | x <- uniqs, (numTimeFound line) > 1] | y <- uniqs, (numTimeFound line) > 2]
  -- trace "hi" (show total)
  return  acc

main = do
  contents <- readFile "./input.txt"
  let input = lines contents
  let results = foldl parseLine [(0,0)] input
  print "fuck"
