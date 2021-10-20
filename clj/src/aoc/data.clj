(ns aoc.data
  (:require [clojure.string :as string]))

(defn read-str [year day]
  (slurp (format "data/%d/d%02d.txt" year day)))

(defn read-lines [year day]
  (string/split-lines (read-str year day)))

(defn read-ints [year day]
  (map #(Integer/parseInt %)
       (read-lines year day)))