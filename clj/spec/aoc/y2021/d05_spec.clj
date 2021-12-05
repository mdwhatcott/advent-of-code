(ns aoc.y2021.d05-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d05 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 5))

(def sample-data
  ["0,9 -> 5,9"
   "8,0 -> 0,8"
   "9,4 -> 3,4"
   "2,2 -> 2,1"
   "7,0 -> 7,4"
   "6,4 -> 2,0"
   "0,9 -> 2,9"
   "3,4 -> 1,4"
   "0,0 -> 8,8"
   "5,5 -> 8,2"])

(describe "2021 Day 5"
  (context "Part 1"
    (it "parses input into point range pairs (omitting diagonals)"
      (should= [{:from [0 9] :to [5 9]}
                #_{:from [8 0] :to [0 8]}
                {:from [9 4] :to [3 4]}
                {:from [2 2] :to [2 1]}
                {:from [7 0] :to [7 4]}
                #_{:from [6 4] :to [2 0]}
                {:from [0 9] :to [2 9]}
                {:from [3 4] :to [1 4]}
                #_{:from [0 0] :to [8 8]}
                #_{:from [5 5] :to [8 2]}]
               (sut/parse-endpoints sample-data false)))

    (it "parses input into point range pairs (keeping diagonals)"
      (should= [{:from [0 9] :to [5 9]}
                {:from [8 0] :to [0 8]}
                {:from [9 4] :to [3 4]}
                {:from [2 2] :to [2 1]}
                {:from [7 0] :to [7 4]}
                {:from [6 4] :to [2 0]}
                {:from [0 9] :to [2 9]}
                {:from [3 4] :to [1 4]}
                {:from [0 0] :to [8 8]}
                {:from [5 5] :to [8 2]}]
               (sut/parse-endpoints sample-data true)))

    (it "explodes range records"
      (should= [[0 0] [0 1] [0 2]] (sut/draw-line {:from [0 0] :to [0 2]}))
      (should= [[0 2] [0 1] [0 0]] (sut/draw-line {:from [0 2] :to [0 0]}))
      (should= [[0 0] [1 1] [2 2]] (sut/draw-line {:from [0 0] :to [2 2]})))

    (it "solves with sample data"
      (should= 5 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 6397 (sut/part1 real-data)))

    )

  (context "Part 2"
    (it "solves with sample data"
      (should= 12 (sut/part2 sample-data)))

    (it "solves with real data"
        (should= 22335 (sut/part2 real-data)))
    )
  )
