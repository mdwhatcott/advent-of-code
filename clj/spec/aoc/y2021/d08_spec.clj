(ns aoc.y2021.d08-spec
  (:require [speclj.core :refer :all]
            [aoc.y2021.d08 :as sut]
            [aoc.data :as data]))

(def real-data
  (data/read-lines 2021 8))

(def one-line-sample
  "acedgfb cdfbe gcdfa fbcad dab cefabd cdfgeb eafb cagedb ab | cdfeb fcadb cdfeb cdbaf")

(def sample-data
  ["be cfbegad cbdgef fgaecd cgeb fdcge agebfd fecdb fabcd edb | fdgacbe cefdb cefbgd gcbe"
   "edbfga begcd cbg gc gcadebf fbgde acbgfd abcde gfcbed gfec | fcgedb cgb dgebacf gc"
   "fgaebd cg bdaec gdafb agbcfd gdcbef bgcad gfac gcb cdgabef | cg cg fdcagb cbg"
   "fbegcd cbd adcefb dageb afcb bc aefdc ecdab fgdeca fcdbega | efabcd cedba gadfec cb"
   "aecbfdg fbg gf bafeg dbefa fcge gcbea fcaegb dgceab fcbdga | gecf egdcabf bgf bfgea"
   "fgeab ca afcebg bdacfeg cfaedg gcfdb baec bfadeg bafgc acf | gebdcfa ecba ca fadegcb"
   "dbcfg fgd bdegcaf fgec aegbdf ecdfab fbedc dacgb gdcebf gf | cefg dcbef fcge gbcadfe"
   "bdfegc cbegaf gecbf dfcage bdacg ed bedf ced adcbefg gebcd | ed bcgafe cdgba cbgef"
   "egadfb cdbfeg cegd fecab cgb gbdefca cg fgcdab egfdb bfceg | gbdfcae bgc cg cgb"
   "gcafb gcf dcaebfg ecagb gf abcdeg gaef cafbge fdbac fegbdc | fgae cfgab fg bagce"])

(describe "2021 Day 8"
  (context "Part 1"
    (it "solves with sample data"
      (should= 26 (sut/count-unique-digits sample-data)))

    (it "solves with real data"
      (should= 519 (sut/count-unique-digits real-data)))

    )

  #_(context "Part 2"
      #_(it "solves with sample data"
          (should= 0 (sut/part2 sample-data)))

      #_(it "solves with real data"
          (should= 0 (sut/part2 real-data)))
      )
  )
