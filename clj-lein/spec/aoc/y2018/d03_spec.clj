(ns aoc.y2018.d03-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d03 :as sut]
            [aoc.data :as data]))

(def example-claim-1 "#232 @ 364,934: 27x10")

(def sample-data
  ["#1 @ 1,3: 4x4"
   "#2 @ 3,1: 4x4"
   "#3 @ 5,5: 2x2"])

(def real-data (data/read-lines 2018 3))

(describe "2018 Day 3"
  (context "Part 1"
    (it "parses claims"
      (should= {:id 232
                :x  364 :y 934
                :w  27, :h 10} (sut/parse-claim example-claim-1)))

    (it "explodes claims into cells covered"
      (should= [[22 2] [23 2] [24 2]
                [22 3] [23 3] [24 3]]
               (sut/explode-cells {:x 22 :y 2 :w 3 :h 2})))

    (it "overlays claims"
      (let [claim1   {:x 2 :y 2 :w 2 :h 2}
            claim2   {:x 3 :y 3 :w 2 :h 2}
            expected {[2 2] 1 [3 2] 1
                      [2 3] 1 [3 3] 2
                      [4 3] 1 [3 4] 1 [4 4] 1}
            actual   (sut/overlay-claims [claim1 claim2])]
        (should= actual expected)))

    (it "solves with sample input"
      (should= 4 (sut/part1 sample-data)))

    (it "solves with real input"
      (should= 111266 (sut/part1 real-data)))
    )

  (context "Part 2"
    (it "solves with sample input"
      (should= 3 (sut/part2 sample-data)))

    (it "solves with real input"
      (should= 266 (sut/part2 real-data)))
    )
  )
