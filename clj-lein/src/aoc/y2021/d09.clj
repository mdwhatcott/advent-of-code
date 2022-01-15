(ns aoc.y2021.d09)

(def high-point (long \9))

(defn digit [c] (- c (long \0)))

(defn at [lines [x y]]
  (-> lines
      (nth y "")
      (nth x high-point)
      long))

(defn neighbors [[x y]]
  [[(+ 0 x) (inc y)]
   [(+ 0 x) (dec y)]
   [(inc x) (+ 0 y)]
   [(dec x) (+ 0 y)]])

(defn low-points [lines]
  (for [x (range (count (first lines)))
        y (range (count lines))
        :let [point     [x y]
              here      (at lines point)
              neighbors (neighbors point)
              neighbors (map (partial at lines) neighbors)]
        :when (= 4 (count (filter #(< here %) neighbors)))] point))

(defn part1 [lines]
  (->> (low-points lines)
       (map (partial at lines))
       (map digit)
       (map inc)
       (apply +)))

(defn measure-basin-BFS [lines low-point]
  (let [value (at lines low-point)]
    (if (>= value high-point)
      0 (loop [size 0 seen #{} queue [low-point]]
          (if (empty? queue)
            (count seen)
            (let [point (peek queue)
                  nexts (neighbors point)
                  nexts (remove #(or (seen point)
                                     (>= (at lines %1) high-point)) nexts)]
              (recur (inc size)
                     (conj seen point)
                     (apply conj (pop queue) nexts))))))))

(defn basins [lines]
  (->> (low-points lines)
       (map (partial measure-basin-BFS lines))))

(defn part2 [lines]
  (->> (basins lines)
       (sort >)
       (take 3)
       (apply *)))
