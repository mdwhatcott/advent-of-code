(ns aoc.y2021.d07-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d07 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-ints 2021 7 #","))

(def sample-data
  [16 1 2 0 4 2 7 1 2 14])

(describe "2021 Day 7"
  (context "Part 1"
    (it "solves with sample data"
      (should= 37 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 328187 (sut/part1 real-data)))

    )

  (context "Part 2"
    (it "solves with sample data"
      (should= 168 (sut/part2 sample-data)))

    (it "solves with real data"
      (should= 91257582 (sut/part2 real-data)))
    )
  )
