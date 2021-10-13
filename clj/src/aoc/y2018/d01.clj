(ns aoc.y2018.d01
  (:require [clojure.string :as string]))

(defn lines->ints [input]
  (->> (string/split-lines input)
       (map #(Integer/parseInt %))))

(defn day01-part1 [input]
  (->> (lines->ints input)
       (apply +)))

(defn day01-part2 [input]
  (let [original (lines->ints input)]
    (loop [at    0
           seen  #{}
           steps (cycle original)]
      (if (contains? seen at)
        at (recur (+ at (first steps))
                  (conj seen at)
                  (rest steps))))))
