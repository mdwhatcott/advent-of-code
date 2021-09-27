(ns aoc.y2020.d06
  (:require [clojure.string :as string]))


(defn count-any-yes [group]
  (->> (string/replace group "\n" "")
       set
       count))

(defn sum-group-yes [input counter]
  (->> (string/split input #"\n\n")
       (map counter)
       (apply +)))

(def actual-input (slurp "src/aoc/y2020/d06.txt"))

(defn part1 [input]
  (sum-group-yes input count-any-yes))