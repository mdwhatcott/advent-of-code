(ns aoc.y2021.d03)

(defn nth-bits [col lines]
  (map #(nth (seq %) col) lines))

(defn by-column-key [comp prefer kvs]
  (let [sorted (sort-by last comp kvs)]
    (if (apply = (vals sorted)) prefer (ffirst sorted))))

(defn bin-not [binary]
  (apply str (map {\1 \0 \0 \1} binary)))

(defn bin->dec [str]
  (Integer/parseInt str 2))

(defn gamma-rate [lines]
  (as-> (range (count (first lines))) $
        (map nth-bits $ (cycle [lines]))
        (map frequencies $)
        (map (partial by-column-key > 0) $)
        (apply str $)))

(defn epsilon-rate [lines]
  (bin-not (gamma-rate lines)))

(defn power-consumption [lines]
  (* (bin->dec (gamma-rate lines))
     (bin->dec (epsilon-rate lines))))

(defn rating-step
  [comparison preference {:keys [column lines] :as state}]
  (let [bits  (nth-bits column lines)
        freq  (frequencies bits)
        keep  (by-column-key comparison preference freq)
        lines (filter #(= keep (nth % column)) lines)]
    (as-> state $
          (update $ :column inc)
          (assoc $ :lines lines))))

(defn filtered-rating [lines comparison preference]
  (let [initial {:column 0 :lines (set lines)}
        step    (partial rating-step comparison preference)]
    (as-> (iterate step initial) $
          (drop-while #(> (count (:lines %)) 1) $)
          (first $)
          (:lines $)
          (first $))))

(defn oxygen-generator-rating [lines]
  (filtered-rating lines > \1))

(defn CO2-scrubber-rating [lines]
  (filtered-rating lines < \0))

(defn life-support-rating [lines]
  (* (bin->dec (oxygen-generator-rating lines))
     (bin->dec (CO2-scrubber-rating lines))))