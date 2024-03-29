(ns aoc.y2019.d02-spec
  (:require
    [speclj.core :refer :all]
    [aoc.y2019.d02 :refer :all]
    [aoc.data :as data]))

(def initial-memory
  ; d02.txt: 3,4,3,5,1,...
  (->> (data/read-str 2019 2)
       (format "[%s]")
       clojure.edn/read-string))

(describe "2019 Day 2"
  (it "solves part 1"
    (should= 3101878 (part1 initial-memory)))

  (it "solves part 2"
    (should= 8444 (part2 initial-memory)))

  )