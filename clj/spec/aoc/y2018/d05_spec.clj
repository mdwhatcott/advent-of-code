(ns aoc.y2018.d05-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d05 :as sut]
            [aoc.data :as data]))

(def sample-input "dabAcCaCBAcCcaDA")
(def real-input (data/read 2018 5))

(describe "2018 Day 5"
  (context "Part 1"
    (it "solves with sample input"
      (should= 10 (sut/part1 sample-input)))

    (it "solves with real input"
      (should= 11264 (sut/part1 real-input)))
    )

  (context "Part 2"
    (it "solves with sample input"
      (should= 4 (sut/part2 sample-input)))

    (it "solves with real input"
      (should= 4552 (sut/part2 real-input)))
    )
  )
