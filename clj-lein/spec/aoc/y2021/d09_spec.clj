(ns aoc.y2021.d09-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d09 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 9))

(def sample-data
  ["2199943210"
   "3987894921"
   "9856789892"
   "8767896789"
   "9899965678"])

(describe "2021 Day 9"
  (context "Part 1"
    (it "solves with sample data"
      (should= 15 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 535 (sut/part1 real-data)))
    )

  (context "Part 2"
    (it "measures a basin"
      (should= 0 (sut/measure-basin-BFS sample-data [2 0]))
      (should= 1 (sut/measure-basin-BFS ["09"
                                     "99"] [0 0]))
      (should= 2 (sut/measure-basin-BFS ["00"
                                     "99"] [0 0]))
      (should= 3 (sut/measure-basin-BFS ["00"
                                     "09"] [0 0])))
    (it "measures basin sizes"
      (should= [3 9 9 14] (sort (sut/basins sample-data))))

    (it "solves with sample data"
        (should= 1134 (sut/part2 sample-data)))

    (it "solves with real data"
        (should= 1122700 (sut/part2 real-data)))
    )
  )
