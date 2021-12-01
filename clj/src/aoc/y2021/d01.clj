(ns aoc.y2021.d01)

(defn count-increases [data]
  (as-> data $
        (partition 2 1 $)
        (map #(- (first %1) (last %1)) $)
        (filter neg? $)
        (count $)))

(defn count-window-increases [window data]
  (as-> data $
        (partition window 1 $)
        (map #(apply + %) $)
        (count-increases $)))
