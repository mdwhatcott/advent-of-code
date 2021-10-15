(ns aoc.y2018.d01-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d01 :as sut]
            [aoc.data :as data]))

(describe "2018 Day 1"
  (context "Part 1"
    (it "solves simple examples"
      (should= 3, (sut/part1 [1 1 1]))
      (should= 0, (sut/part1 [1 1 -2]))
      (should= -6 (sut/part1 [-1 -2 -3])))

    (it "solves with real input"
      (should= 406 (sut/part1 (data/read-ints 1)))))

  (context "Part 2"
    (it "solves simple examples"
      (should= 0, (sut/part2-loop [1 -1]))
      (should= 10 (sut/part2-loop [3 3 4 -2 -4]))
      (should= 5, (sut/part2-loop [-6 3 8 5 -6]))
      (should= 14 (sut/part2-loop [7 7 2 -7 -4])))

    (it "solves with real input"
      (should= 312 (sut/part2-loop (data/read-ints 1))))
    ; Returned to 312 after cycling through 139324 input numbers,
    ; of which there were 1014

    (it "solves with reductions"
      (should= 312 (sut/part2-reductions (data/read-ints 1))))
    )
  )
