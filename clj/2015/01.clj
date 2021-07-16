(def input (slurp "01.txt"))

(println "part 1:"
         (- (count (filter #(= \( %) input))
            (count (filter #(= \) %) input))))
