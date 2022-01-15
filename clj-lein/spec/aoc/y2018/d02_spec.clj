(ns aoc.y2018.d02-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d02 :as sut]
            [aoc.data :as data]))

(def examples-part1 ["abcdef"
                     "bababc"
                     "abbcde"
                     "abcccd"
                     "aabcdd"
                     "abcdee"
                     "ababab"])

(def examples-part2 ["abcde"
                     "fghij"
                     "klmno"
                     "pqrst"
                     "fguij"
                     "axcye"
                     "wvxyz"])

(def real-inputs (data/read-lines 2018 2))

(describe "2018 Day 2"
  (context "Part 1"
    (it "solves simple examples"
      (should-not (sut/has-n-repeats? 2 "abcde"))
      (should (sut/has-n-repeats? 2 "bababc"))
      (should (sut/has-n-repeats? 3 "bababc")))

    (it "counts all repeats of n in a list"
      (should= 4 (sut/count-repeats 2 examples-part1))
      (should= 3 (sut/count-repeats 3 examples-part1)))

    (it "solves with real input"
      (should= 5976 (sut/part1 real-inputs))))

  (context "Part 2"
    (it "analyzes two strings"
      (should= 1 (sut/diff-count "fghij"
                                 "fguij"))
      (should= 5 (sut/diff-count "abcde"
                                 "fguij"))
      (should= 5 (sut/diff-count "xrecqmdonskvzupalfkwhjctdb"
                                 "xrlgqmavnskvzupalfiwhjctdb")))

    (it "solves simple examples"
      (should= "fgij" (sut/part2 examples-part2)))

    (it "solves with real input"
        (should= "xretqmmonskvzupalfiwhcfdb"
                 (sut/part2 real-inputs)))

    )
  )
