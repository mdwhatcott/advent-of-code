(ns aoc.y2018.d14)

(def input 290431)

(def seed {:scores [3 7] :elf1 0 :elf2 1})

(defn make-recipes [{:keys [scores elf1 elf2]}]
  (let [recipe1 (nth scores elf1)
        recipe2 (nth scores elf2)
        sum     (+ recipe1 recipe2)
        score1  (quot sum 10)
        score2  (mod sum 10)
        scores  (if (zero? score1) scores (conj scores score1))
        scores  (conj scores score2)
        elf1    (mod (inc (+ elf1 recipe1)) (count scores))
        elf2    (mod (inc (+ elf2 recipe2)) (count scores))]
    (when (zero? (mod (count scores) 100))
      (println (count scores)))
    {:scores scores :elf1 elf1 :elf2 elf2}))

(defn ten-scores-after-n-iterations [n]
  (let [ceiling (+ n 10)]
    (as-> (iterate make-recipes seed) $
          (drop-while #(< (count (:scores %)) ceiling) $)
          (first $)
          (:scores $)
          (drop (- (count $) 10) $)
          (apply str $))))

(defn last-n [n s]
  (loop [s s result []]
    (if (or (empty? s) (= n (count result)))
      (vec (reverse result))
      (recur (pop s) (conj result (peek s))))))

(defn find-suffix [suffix]
  (as-> (iterate make-recipes seed) $
        (drop-while #(not= suffix (last-n (count suffix) (:scores %))) $)
        (first $)
        (:scores $)
        (- (count $) (count suffix))))
