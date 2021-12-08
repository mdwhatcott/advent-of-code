(ns aoc.y2021.d08
  (:require [aoc.data :as data]
            [clojure.string :as string]
            [clojure.pprint :as pprint]))

(defn part1 [data])

(defn part2 [data])

(defn count-unique-digits [lines]
  (as-> lines $
        (map #(string/split % #"\s\|\s") $)
        (map last $)
        (string/join " " $)
        (string/split $ #"\s")
        (map count $)
        (frequencies $)
        (+ (get $ 2)                                        ; 1
           (get $ 4)                                        ; 4
           (get $ 3)                                        ; 7
           (get $ 7))))                                     ; 8
