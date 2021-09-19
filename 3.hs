ownLength :: Num p => [a] -> p
ownLength [] = 0
ownLength (_ : xs) = 1 + ownLength xs

getTreeCount :: (Show t, Num t) => [[Char]] -> Int -> Int -> Int -> Int -> Int -> t -> IO ()
getTreeCount sample index size step down startIdx treeCount = do
  if (index + 1) >= size
    then do
      putStrLn ("slope (step: " ++ show step ++ ", down: " ++ show down ++ ") " ++ "tree count :>> " ++ show treeCount)
    else do
      if sample !! (index + down) !! (startIdx) == '#'
        then do
          let updatedStartIdx = (index + step) `mod` ownLength (sample !! index)
          getTreeCount sample (index + down) size step down updatedStartIdx (treeCount + 1)
        else do
          getTreeCount sample (index + down) size step down startIdx treeCount

getTotalTreesTraversedBySlope :: [[Char]] -> Int -> Int -> IO ()
getTotalTreesTraversedBySlope sample step down = do
  let size = ownLength sample
  getTreeCount sample 0 size step down 0 0

main :: IO ()
main = do
  let sample = ["#ab", "#ab", "abc"]
  getTotalTreesTraversedBySlope sample 1 1
  getTotalTreesTraversedBySlope sample 3 1
  getTotalTreesTraversedBySlope sample 5 1
  getTotalTreesTraversedBySlope sample 7 1
  getTotalTreesTraversedBySlope sample 1 2
