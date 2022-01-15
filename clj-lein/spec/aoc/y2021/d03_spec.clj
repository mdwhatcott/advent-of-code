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
      (should= "10110" (sut/gamma-rate sample-data)))

    (it "solves with sample data"
      (should= 198 (sut/power-consumption sample-data)))

    (it "solves with real data"
      (should= 749376 (sut/power-consumption real-data)))

    )

  (context "Part 2"
    (it "steps to find oxygen generator rating"
      (let [data  (set sample-data)
            state {:column 0 :lines data}]
        (should=, {:column 1 :lines (remove #(= \0 (first %)) data)}
                  (sut/rating-step > \1 state))))

    (it "finds the oxygen generator rating"
      (should= "10111" (sut/oxygen-generator-rating sample-data)))

    (it "finds the CO2 scrubber rating"
      (should= "01010" (sut/CO2-scrubber-rating sample-data)))

    (it "solves with sample data"
      (should= 230 (sut/life-support-rating sample-data)))

    (it "solves with real data"
      (should= 2372923 (sut/life-support-rating real-data)))
    )
  )
