(ns aoc.y2018.d01)

(defn part1 [ints]
  (apply + ints))

(defn until-seen-in [seen x]
  (if (seen x)
    (reduced x)
    (conj seen x)))

(defn part2-reductions [ints]
  (as-> (cycle ints) $
        (reductions + $)
        (reduce until-seen-in #{0} $)))
