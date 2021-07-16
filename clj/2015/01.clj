(def input (slurp "01.txt"))

(defn count-char [char input]
  (count (filter #(= char %) input)))

(defn walk [input]
  (- (count-char \( input)
     (count-char \) input)))

(println "part 1:" (walk input))

;; super inefficient!
;; (walks incrementally larger sub-
;; strings until we arrive at -1)
(defn walk-until [low-point input]
  (->> (range 1 (inc (count input)))
       (map #(subs input 0 %))
       (map walk)
       (take-while #(> % low-point))
       (count)
       (inc)))

(println "part 2 (slow):" (walk-until -1 input))

(def up \()
(defn step [char]
  (if (= char up) 1 -1))

(defn part2 [input cursor floor]
  (if (= floor -1) cursor
    (let [next-input  (drop 1 input)
          next-cursor (inc cursor)
          next-floor  (+ floor (step (first input)))]
      (part2 next-input next-cursor next-floor))))

(println "part 2 (fast):" (part2 (seq input) 0 0))



