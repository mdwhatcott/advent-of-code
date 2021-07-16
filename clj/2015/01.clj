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

(println "part 2:" (walk-until -1 input))
