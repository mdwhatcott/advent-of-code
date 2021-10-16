(ns aoc.y2018.d03)

(defn parse-claim [raw]
  ; Sample Claim: #266 @ 105,418: 27x14
  (let [matches (re-matches #"#(\d+) @ (\d+),(\d+): (\d+)x(\d+)" raw)
        [id x1 y1 w h] (map #(Integer/parseInt %) (rest matches))]
    {:id id :x x1 :y y1 :w w :h h}))

(defn explode-cells [{:keys [x y w h]}]
  (for [y (range y (+ y h))
        x (range x (+ x w))] [x y]))

(defn overlay-claims [claims]
  (frequencies (mapcat explode-cells claims)))

(defn overlapping [claims]
  (->> (overlay-claims claims)
       (remove #(= 1 (second %)))))

(defn part1 [lines]
  (->> lines (map parse-claim) overlapping count))

(defn part2 [lines]
  (let [claims (map parse-claim lines)
        counts (overlay-claims claims)]
    (first
      (for [claim claims
            :let [cells       (explode-cells claim)
                  cell-counts (for [cell cells] (get counts cell))
                  ones        (filter #(= 1 %) cell-counts)]
            :when (= (count cells) (count ones))]
        (:id claim)))))
