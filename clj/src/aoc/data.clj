(ns aoc.data
  (:require [clojure.string :as string]))

(defn str->int [s] (Integer/parseInt s))

(defn read-str [year day]
  (slurp (format "data/%d/d%02d.txt" year day)))

(defn read-lines [year day]
  (string/split-lines (read-str year day)))

(defn read-words [year day]
  (string/split (read-str year day) #"\s+"))

(defn read-ints [year day]
  (map str->int (read-words year day)))