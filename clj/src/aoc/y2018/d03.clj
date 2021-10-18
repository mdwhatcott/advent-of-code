(ns aoc.y2018.d03)

(defn parse-claim [raw]
  ; Sample Claim: #266 @ 105,418: 27x14
  (let [matches (re-seq #"\d+" raw)
        [id x1 y1 w h] (map #(Integer/parseInt %) matches)]
    {:id id :x x1 :y y1 :w w :h h}))

(defn explode-cells [{:keys [x y w h]}]
  (for [y (range y (+ y h))
        x (range x (+ x w))] [x y]))

(defn overlay-claims [claims]
  (frequencies (mapcat explode-cells claims)))

(defn part1 [lines]
  (->> (map parse-claim lines)
       overlay-claims
       (remove #(= 1 (second %)))
       count))

(defn part2 [lines]
  (let [claims (map parse-claim lines)
        counts (overlay-claims claims)]
    (first
      (for [claim claims
            :let [cells   (explode-cells claim)
                  counted (for [cell cells] (counts cell))
                  ones    (filter #(= 1 %) counted)]
            :when (= (count cells) (count ones))]
        (:id claim)))))
