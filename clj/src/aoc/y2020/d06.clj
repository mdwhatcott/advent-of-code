(ns aoc.y2020.d06
  (:require [clojure.string :as string]
            [aoc.data :as data]))


(defn count-any-yes [group]
  (->> (string/replace group "\n" "")
       set
       count))

(defn sum-group-yes [input counter]
  (->> (string/split input #"\n\n")
       (map counter)
       (apply +)))

(def actual-input (data/read 2020 6))

(defn part1 [input]
  (sum-group-yes input count-any-yes))