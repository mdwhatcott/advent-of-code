(ns aoc.data
  (:require [clojure.string :as string]))

(defn read [day]
  (slurp (format "spec/aoc/y2018/d%02d.txt" day)))

(defn read-lines [day]
  (string/split-lines (read day)))

(defn read-ints [day]
  (map #(Integer/parseInt %) (read-lines day)))