(ns aoc.y2019.intcode)

(defn op4 [f {:keys [pointer memory] :as state}]
  (let [[a b c] (take 3 (drop (inc pointer) memory))
        noun (nth memory a)
        verb (nth memory b)
        out  (f noun verb)]
    (assoc state :memory (assoc memory c out)
                 :pointer (+ 4 pointer))))

(defn halt [state]
  (assoc state :running false
               :pointer (inc (:pointer state))))

(defn input [{:keys [inputs pointer memory] :as state}]
  (if (empty? inputs)
    state
    (let [address (nth memory (inc pointer))]
      (-> state
          (assoc-in [:memory address] (first inputs))
          (assoc :pointer (+ 2 pointer))
          (assoc :inputs (rest inputs))))))

(defn output [{:keys [pointer memory outputs] :as state}]
  (let [address (nth memory (inc pointer))]
    (assoc state :outputs (conj outputs (nth memory address))
                 :pointer (+ 2 pointer))))

(defn tick [{:keys [pointer memory] :as state}]
  (cond
    (not (contains? state :running)) state
    (not (:running state)) (dissoc state :running)
    :else (let [opcode (nth memory pointer 0)]
            (case opcode
              1 (op4 + state)
              2 (op4 * state)
              3 (input state)
              4 (output state)
              99 (halt state)))))

(defn run [initial-state]
  (let [defaults {:pointer 0
                  :inputs  []
                  :running true}]
    (->> (merge defaults initial-state)
         (iterate tick)
         (take-while #(contains? % :running))
         last)))

(defn run-simple [memory]
  (let [state {:memory memory}]
    (:memory (run state))))
