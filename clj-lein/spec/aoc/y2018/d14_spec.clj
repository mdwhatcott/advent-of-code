(ns aoc.y2018.d14-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d14 :as sut]))

(describe "2018 Day 14"
  (context "Part 1"
    (it "solves a single iteration"
      (let [output (sut/make-recipes sut/seed)]
        (should= {:scores [3 7 1 0]
                  :elf1   0
                  :elf2   1} output)))

    (it "solves 9 iterations"
      (should= "5158916779" (sut/ten-scores-after-n-iterations 9)))

    (it "solves with real data"
      (should= "1776718175" (sut/ten-scores-after-n-iterations sut/input)))

    )

  (context "Part 2"
    (it "solves with sample data"
      (should= 9 (sut/find-suffix "51589"))
      (should= 5 (sut/find-suffix "01245"))
      (should= 18 (sut/find-suffix "92510"))
      (should= 2018 (sut/find-suffix "59414")))

    #_(it "solves with real data (albiet slowly - 13s)"
      (should= 20220949 (sut/find-suffix "290431")))
    )
  )
