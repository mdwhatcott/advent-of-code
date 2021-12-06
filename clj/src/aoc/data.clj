(ns aoc.data
  (:require [clojure.string :as string]))

(defn str->int [s] (Integer/parseInt s))

(defn read-str [year day]
  (slurp (format "data/%d/d%02d.txt" year day)))

(defn read-lines [year day]
  (string/split-lines (read-str year day)))

(defn read-words
  ([year day]
   (read-words year day #"\s+"))
  ([year day re-sep]
   (string/split (read-str year day) re-sep)))

(defn read-ints
  ([year day]
   (read-ints year day #"\s+"))
  ([year day re-sep]
   (map str->int (read-words year day re-sep))))