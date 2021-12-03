(ns aoc.y2021.d03
  (:require [clojure.string :as string]
            [aoc.data :as data]))

(defn column-values [col lines]
  (map #(nth (seq %) col) lines))

(defn max-key [kvs]
  (first (last (sort-by last kvs))))

(defn most-common-by-column [lines]
  (as-> (range (count (first lines))) $
        (map column-values $ (cycle [lines]))
        (map frequencies $)
        (map max-key $)
        (apply str $)))

(defn invert-binary [binary]
  (apply str (map {\1 \0 \0 \1} binary)))

(defn part1 [data]
  (let [gamma   (most-common-by-column data)
        epsilon (invert-binary gamma)]
    (* (Integer/parseInt gamma 2)
       (Integer/parseInt epsilon 2))))
