(ns aoc.y2018.d14
  (:require
    [benchmarks.bench :as bench]
    [clojure.string :as str]))

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
    {:scores scores :elf1 elf1 :elf2 elf2}))

(defn ten-scores-after-n-iterations [n]
  (let [ceiling (+ n 10)]
    (as-> (iterate make-recipes seed) $
          (drop-while #(< (count (:scores %)) ceiling) $)
          (first $)
          (:scores $)
          (drop (- (count $) 10) $)
          (apply str $))))

(defn iterate-n [seed]
  (as-> (iterate make-recipes seed) $
        (drop (:iterate-n seed) $)
        (first $)
        (do
          (when (>= (:iterate-n seed) 4096)
            (println (:iterate-n seed))) $)
        (assoc $ :iterate-n (* 2 (:iterate-n seed)))))

(defn suffix-absent? [suffix state]
  (nil? (str/index-of (apply str (:scores state)) suffix)))

(defn find-suffix [suffix]
  (as-> (assoc seed :iterate-n 1024) $
        (iterate (partial iterate-n) $)
        (drop-while (partial suffix-absent? suffix) $)
        (first $)
        (:scores $)
        (apply str $)
        (str/index-of $ suffix)))

(defn -main []
  (println (bench/report #(find-suffix (str input)))))