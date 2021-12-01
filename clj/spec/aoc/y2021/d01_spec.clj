(ns aoc.y2021.d01-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d01 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-ints 2021 1))

(def sample-data
  [199
   200
   208
   210
   200
   207
   240
   269
   260
   263])

(describe "2021 Day 1"
  (context "Part 1"
    (it "solves with sample data"
      (should= 7 (sut/count-increases sample-data)))

    (it "solves with real data"
      (should= 1688 (sut/count-increases real-data)))

    )

  (context "Part 2"
      (it "solves with sample data"
        (should= 5 (sut/count-window-increases 3 sample-data)))

      (it "solves with real data"
        (should= 1728 (sut/count-window-increases 3 real-data)))
    )
  )
