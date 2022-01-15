(ns aoc.y2021.d02-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d02 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 2))

(def sample-data
  ["forward 5"
   "down 5"
   "forward 8"
   "up 3"
   "down 8"
   "forward 2"])

(describe "2021 Day 2"
  (context "Part 1"
    (it "parses lines"
      (should= {:action "forward" :n 5} (sut/parse-line "forward 5")))

    (it "solves with sample data"
      (should= {:horizontal 15 :depth 10} (sut/traverse-exact sample-data))
      (should= 150 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 2150351 (sut/part1 real-data)))

    )

  (context "Part 2"
    (it "makes a single step"
      (should= {:horizontal 1 :depth 0 :aim 0}
               (sut/step-aim "forward 1" {:horizontal 0 :depth 0 :aim 0}))

      (should= {:horizontal 0 :depth 0 :aim 5}
               (sut/step-aim "down 5" {:horizontal 0 :depth 0 :aim 0}))

      (should= {:horizontal 0 :depth 0 :aim 2}
               (sut/step-aim "up 3" {:horizontal 0 :depth 0 :aim 5}))

      (should= {:horizontal 13 :depth 40 :aim 5}
               (sut/step-aim "forward 8" {:horizontal 5 :depth 0 :aim 5})))

    (it "solves with sample data"
      (should= {:horizontal 15 :depth 60} (sut/traverse-aim sample-data))
      (should= 900 (sut/part2 sample-data)))

    (it "solves with real data"
      (should= 1842742223 (sut/part2 real-data)))
    )
  )
