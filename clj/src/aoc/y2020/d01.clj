(ns aoc.y2020.d01
  (:require
    [clojure.edn :as edn]
    [clojure.string :as string]
    [aoc.data :as data]))

(def actual-input (data/read-str 2020 1))

(defn part1 [input]
  (as-> input $
        (string/split $ #"\s")
        (map edn/read-string $)
        (for [a (range 0 (count $))
              b (range a (count $))
              :let [A (nth $ a)
                    B (nth $ b)]
              :when (= (+ A B) 2020)] (* A B))
        (first $)))

(defn part2 [input]
  (as-> input $
        (string/split $ #"\s")
        (map edn/read-string $)
        (for [a (range 0 (count $))
              b (range a (count $))
              c (range b (count $))
              :let [A (nth $ a)
                    B (nth $ b)
                    C (nth $ c)]
              :when (= (+ A B C) 2020)] (* A B C))
        (first $)))
