(ns aoc.data
  (:require [clojure.string :as string]))

(defn read [year day]
  (slurp (format "spec/aoc/y%d/d%02d.txt" year day)))

(defn read-lines [year day]
  (string/split-lines (read year day)))

(defn read-ints [year day]
  (map #(Integer/parseInt %)
       (read-lines year day)))