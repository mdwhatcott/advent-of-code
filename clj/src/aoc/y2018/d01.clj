(ns aoc.y2018.d01
  (:require [clojure.string :as string]))

(defn lines->ints [input]
  (->> (string/split-lines input)
       (map #(Integer/parseInt %))))

(defn part1 [input]
  (->> (lines->ints input)
       (apply +)))

(defn part2-loop [input]
  (let [original (lines->ints input)]
    (loop [at    0
           seen  #{}
           steps (cycle original)]
      (if (contains? seen at)
        at (recur (+ at (first steps))
                  (conj seen at)
                  (rest steps))))))

(defn until-seen-in [seen x]
  (if (seen x)
    (reduced x)
    (conj seen x)))

(defn part2-reductions [input]
  (as-> (cycle (lines->ints input)) $
        (reductions + $)
        (reduce until-seen-in #{0} $)))
