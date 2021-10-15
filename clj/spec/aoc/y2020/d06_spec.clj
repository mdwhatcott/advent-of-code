(ns aoc.y2020.d06-spec
  (:require [speclj.core :refer :all]
            [aoc.y2020.d06 :refer :all]))

(def example-input-1
  (str "abcx\n"
       "abcy\n"
       "abcz\n"))

(def example-input-2
  (str "abc\n"
       "\n"
       "a\n"
       "b\n"
       "c\n"
       "\n"
       "ab\n"
       "ac\n"
       "\n"
       "a\n"
       "a\n"
       "a\n"
       "a\n"
       "\n"
       "b"))

(describe "2020 Day 6"
  (it "counts distinct items per group"
    (should= 6 (count-any-yes example-input-1)))

  (it "sums each groups' distinct items"
    (should= 11 (sum-group-yes example-input-2 count-any-yes)))

  (it "solves part 1"
    (should= 6775 (part1 actual-input)))


  )
