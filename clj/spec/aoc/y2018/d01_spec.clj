(ns aoc.y2018.d01-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d01 :as sut]
            [clojure.string :as string]))

(defn convert-example [input]
  (string/replace input ", " "\n"))

(def real-input (slurp "spec/aoc/y2018/d01.txt"))

(describe "2018 Day 1"
  (context "Part 1"
    (it "solves simple examples"
      (should= 3, (sut/part1 (convert-example "+1, +1, +1")))
      (should= 0, (sut/part1 (convert-example "+1, +1, -2")))
      (should= -6 (sut/part1 (convert-example "-1, -2, -3"))))

    (it "solves with real input"
      (should= 406 (sut/part1 real-input))))

  (context "Part 2"
    (it "solves simple examples"
      (should= 0, (sut/part2-loop (convert-example "+1, -1")))
      (should= 10 (sut/part2-loop (convert-example "+3, +3, +4, -2, -4")))
      (should= 5, (sut/part2-loop (convert-example "-6, +3, +8, +5, -6")))
      (should= 14 (sut/part2-loop (convert-example "+7, +7, -2, -7, -4"))))

    (it "solves with real input"
      (should= 312 (sut/part2-loop real-input)))
    ; Returned to 312 after cycling through 139324 input numbers,
    ; of which there were 1014

    (it "solves with reductions"
      (should= 312 (sut/part2-reductions real-input)))
    )
  )