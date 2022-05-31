(ns y2016.d03
  (:require [aoc.aoc :as aoc]
            [clojure.string :as str]
            [clojure.test :refer [deftest is]]))

(defn valid-triangle? [a b c]
  (and (> (+ a b) c)
       (> (+ b c) a)
       (> (+ a c) b)))

(defn part1 [data]
  (as-> data $
        (str/split $ #"\s")
        (remove #(= % "") $)
        (map #(Integer/parseInt %) $)
        (partition 3 $)
        (map #(apply valid-triangle? %) $)
        (filter true? $)
        (count $)))

(defn triangles-in-columns
  [[a d g
    b e h
    c f i]]
  [[a b c]
   [d e f]
   [g h i]])

(defn part2 [data]
  (as-> data $
        (str/split $ #"\s")
        (remove #(= % "") $)
        (map #(Integer/parseInt %) $)
        (partition 9 $)
        (mapcat triangles-in-columns $)
        (map #(apply valid-triangle? %) $)
        (filter true? $)
        (count $)))

(deftest part-1
  (is (= 983 (part1 (aoc/input-string 2016 3)))))

(deftest part-2
  (is (= 1836 (part2 (aoc/input-string 2016 3)))))

(aoc/run-tests)
