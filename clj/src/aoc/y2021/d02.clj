(ns aoc.y2021.d02
  (:require [aoc.data :as data]
            [clojure.string :as string]))

(defn parse-line [line]
  (let [words (string/split line #"\s")]
    {:action (first words)
     :n      (data/str->int (last words))}))

(defn sum-actions [steps action]
  (->> steps
       (filter #(= action (:action %)))
       (map :n)
       (apply +)))

(defn traverse-exact [data]
  (let [steps    (map parse-line data)
        forwards (sum-actions steps "forward")
        ups      (sum-actions steps "up")
        downs    (sum-actions steps "down")]
    {:horizontal forwards
     :depth      (- downs ups)}))

(defn travel [data diver]
  (apply * (vals (diver data))))

(defn part1 [data]
  (travel data traverse-exact))

(defn step-aim [line state]
  (let [parsed (parse-line line)]
    (case (:action parsed)
      "up",,,,, (update state :aim - (:n parsed))
      "down",,, (update state :aim + (:n parsed))
      "forward" (-> state
                    (update :horizontal + (:n parsed))
                    (update :depth + (* (:aim state) (:n parsed)))))))

(defn wrap-step-aim [state]
  (update (step-aim (first (:steps state)) state) :steps rest))

(defn traverse-aim [data]
  (let [initial {:steps      data
                 :horizontal 0
                 :depth      0
                 :aim        0}]
    (as-> initial $
          (iterate wrap-step-aim $)
          (drop (count data) $)
          (first $)
          (dissoc $ :steps :aim))))

(defn part2 [data]
  (travel data traverse-aim))
