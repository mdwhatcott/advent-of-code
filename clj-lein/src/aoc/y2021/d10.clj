(ns aoc.y2021.d10)

(def closing? #{\], \), \}, \>})
(def open {\] \[, \) \(, \} \{, \> \<})
(def close (zipmap (vals open) (keys open)))
(def penalty {\) 3, \] 57, \} 1197, \> 25137})
(def boost {\) 1, \] 2, \} 3, \> 4})

(defn analyze [line]
  (loop [stack [] chars (seq line)]
    (if (empty? chars)
      stack
      (let [char         (first chars)
            valid-close? (and (closing? char)
                              (= (open char) (peek stack)))]
        (cond valid-close? (recur (pop stack) (rest chars))
              (closing? char) chars
              :else (recur (conj stack char) (rest chars)))))))

(defn valid? [analysis]
  (not (closing? (first analysis))))

(defn part1 [lines]
  (->> lines
       (map analyze)
       (remove valid?)
       (map first)
       (map penalty)
       (apply +)))

(defn completion [line]
  (loop [closed []
         opened (vec (analyze line))]
    (if (empty? opened)
      closed
      (recur (conj closed (close (peek opened)))
             (vec (take (dec (count opened)) opened))))))

(defn score-completion [completion]
  (loop [score 0 chars (seq completion)]
    (if (empty? chars)
      score
      (let [char (first chars)]
        (recur (+ (* score 5) (boost char))
               (rest chars))))))

(defn part2 [lines]
  (as-> lines $
        (map analyze $)
        (filter valid? $)
        (map completion $)
        (map score-completion $)
        (sort $)
        (drop (/ (dec (count $)) 2) $)
        (first $)))
