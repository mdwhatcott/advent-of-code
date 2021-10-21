(ns aoc.y2018.d06
  (:require [clojure.string :as string]
            [aoc.data :as data]))

(defn parse-coordinates [input]
  (as-> input $
        (string/replace $ "," "")
        (string/split $ #"\s+")
        (map data/str->int $)
        (partition 2 $)
        (set $)))

(defn scope [coords]
  (let [xs    (map first coords)
        ys    (map second coords)
        max-x (apply max xs)
        min-x (apply min xs)
        max-y (apply max ys)
        min-y (apply min ys)]
    [[min-x min-y]
     [max-x max-y]]))

(defn manhattan [c1 c2]
  (let [x1    (first c1)
        x2    (first c2)
        y1    (second c1)
        y2    (second c2)
        max-x (max x1 x2)
        min-x (min x1 x2)
        max-y (max y1 y2)
        min-y (min y1 y2)]
    (+ (- max-x min-x)
       (- max-y min-y))))

(defn arena [[[x1 y1] [x2 y2]]]
  (for [y (range y1 (inc y2))
        x (range x1 (inc x2))] [x y]))

(defn infinite? [landmarks [x y]]
  (let [landmarks (remove #{[x y]} landmarks)
        xs        (map first landmarks)
        ys        (map second landmarks)
        $<x       (partial < x)
        $>x       (partial > x)
        $<y       (partial < y)
        $>y       (partial > y)]
    (or (< (count (filter $<x xs)) 2)
        (< (count (filter $>x xs)) 2)
        (< (count (filter $<y ys)) 2)
        (< (count (filter $>y ys)) 2))))

(defn distances [landmarks location]
  (for [landmark landmarks]
    [(manhattan landmark location) landmark]))

(defn closest [landmarks location]
  (let [distances (sort-by first (distances landmarks location))
        one       (first distances)
        two       (second distances)]
    (if (= (first one) (first two))
      nil
      one)))

(defn calculate-areas [arena landmarks]
  (as-> (map (partial closest landmarks) arena) $
        (remove nil? $)
        (map last $)
        (remove (partial infinite? landmarks) $)
        (frequencies $)))

(defn part1 [input]
  (let [landmarks (parse-coordinates input)
        arena     (arena (scope landmarks))
        areas     (calculate-areas arena landmarks)]
    (apply max (vals areas))))

(defn combined-distance-to [landmarks location]
  (->> landmarks
       (map (partial manhattan location))
       (apply +)))

(defn part2 [n input]
  (let [landmarks (parse-coordinates input)]
    (as-> (arena (scope landmarks)) $
          (map (partial combined-distance-to landmarks) $)
          (remove (partial <= n) $)
          (count $))))
