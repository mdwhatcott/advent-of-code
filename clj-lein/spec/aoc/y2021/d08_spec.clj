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

;   0:      1:      2:      3:      4:
;  aaaa    ....    aaaa    aaaa    ....
; b    c  .    c  .    c  .    c  b    c
; b    c  .    c  .    c  .    c  b    c
;  ....    ....    dddd    dddd    dddd
; e    f  .    f  e    .  .    f  .    f
; e    f  .    f  e    .  .    f  .    f
;  gggg    ....    gggg    gggg    ....
;
;   5:      6:      7:      8:      9:
;  aaaa    aaaa    aaaa    aaaa    aaaa
; b    .  b    .  .    c  b    c  b    c
; b    .  b    .  .    c  b    c  b    c
;  dddd    dddd    ....    dddd    dddd
; .    f  e    f  .    f  e    f  .    f
; .    f  e    f  .    f  e    f  .    f
;  gggg    gggg    ....    gggg    gggg

; acedgfb: 8
; cdfbe: 5
; gcdfa: 2
; fbcad: 3
; dab: 7
; cefabd: 9
; cdfgeb: 6
; eafb: 4
; cagedb: 0
; ab: 1

; cdfeb: 5
; fcadb: 3
; cdfeb: 5
; cdbaf: 3

(describe "2021 Day 8"

  (context "Part 1"
    (it "solves with sample data"
      (should= 26 (sut/part1 sample-data)))

    (it "solves with real data"
      (should= 519 (sut/part1 real-data)))

    )

  (context "Part 2"

    (it "deciphers digits"
      (should= {(set "ab")      1
                (set "dab")     7
                (set "eafb")    4
                (set "acedgfb") 8
                (set "cefabd")  9
                (set "fbcad")   3
                (set "cdfbe")   5
                (set "gcdfa")   2
                (set "cdfgeb")  6
                (set "cagedb")  0} (sut/decipher-digits one-line-sample)))

    (it "reveals the 4-digit code of a line"
      (should= 5353 (sut/reveal-code one-line-sample)))

    (it "solves with sample data"
        (should= 61229 (sut/part2 sample-data)))

    (it "solves with real data"
        (should= 1027483 (sut/part2 real-data)))
    )
  )
