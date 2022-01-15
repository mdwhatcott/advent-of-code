(ns aoc.y2020.d01_spec
  (:require
    [speclj.core :refer :all]
    [aoc.y2020.d01 :refer :all]))

(def example-input
  (str "1721" "\n"
       "979" "\n"
       "366" "\n"
       "299" "\n"
       "675" "\n"
       "1456" "\n"))

(describe "2020 Day 01"
  (it "finds the expected part 1 answer with example input"
    (should= 514579 (part1 example-input))
    (should= 1015476 (part1 actual-input))
    )

  (it "finds the expected part 2 answer with example input"
    (should= 241861950 (part2 example-input))
    (should= 200878544 (part2 actual-input))
    )
  )
