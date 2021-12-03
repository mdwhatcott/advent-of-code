(ns aoc.y2021.d03-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d03 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 3))

(def sample-data
  ["00100"
   "11110"
   "10110"
   "10111"
   "10101"
   "01111"
   "00111"
   "11100"
   "10000"
   "11001"
   "00010"
   "01010"])

(describe "2021 Day 3"
  (context "Part 1"
    (it "picks most common bit value in a column"
      (should= "10110" (sut/most-common-by-column sample-data)))

    (it "solves with sample data"
      (should= 198 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 749376 (sut/part1 real-data)))

    )

  #_(context "Part 2"
      #_(it "solves with sample data")

      #_(it "solves with real data")
    )
  )
