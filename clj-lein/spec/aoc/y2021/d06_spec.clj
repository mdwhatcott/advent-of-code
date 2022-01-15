(ns aoc.y2021.d06-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d06 :as sut]
            [aoc.data :as data]))

(def real-data (data/read-ints 2021 6 #","))
(def sample-data [3 4 3 1 2])

(describe "2021 Day 6"
  (context "Part 1"
    (it "solves with sample data"
      (should= 26 (sut/part1 sample-data 18))
      (should= 5934 (sut/part1 sample-data 80)))

    (it "solves with real data"
        (should= 362666 (sut/part1 real-data 80)))
    )

  (context "Part 2"
      (it "solves with sample data"
          (should= 26984457539 (sut/part2 sample-data)))

      (it "solves with real data"
          (should= 1640526601595 (sut/part2 real-data)))
    )
  )
