(ns aoc.y2019.intcode)

(defn op4 [f {:keys [pointer memory] :as state}]
  (let [[a b c] (take 3 (drop (inc pointer) memory))
        noun (nth memory a)
        verb (nth memory b)
        out  (f noun verb)]
    (assoc state :memory (assoc memory c out)
                 :pointer (+ 4 pointer))))

(defn halt [state]
  (-> state
      (update :pointer inc)
      (assoc :halted true)))

(defn tick [{:keys [pointer memory] :as state}]
  (let [opcode (nth memory pointer)]
    (case opcode
      1 (op4 + state)
      2 (op4 * state)
      99 (halt state))))

(defn run [memory]
  (let [state {:pointer 0
               :memory  memory
               :halted  false}]
    (->> (iterate tick state)
         (take-while #(not (:halted %)))
         last
         :memory)))
