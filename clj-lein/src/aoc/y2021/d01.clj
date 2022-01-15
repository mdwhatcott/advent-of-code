(ns aoc.y2021.d01)

(defn count-increases [data]
  (->> (partition 2 1 data)
       (filter #(apply < %))
       count))

(defn count-window-increases [window data]
  (->> (partition window 1 data)
       (map #(apply + %))
       count-increases))
