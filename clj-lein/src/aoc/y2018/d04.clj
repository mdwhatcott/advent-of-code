(ns aoc.y2018.d04
  (:require [clojure.string :as string]
            [aoc.data :as data]))

; regex matching
; reading 'staggered' records (a guard's shift is spread across multiple lines)
; merge-with concat
; frequencies
; sort-by :map-key

(defn parse-int [input]
  (if (nil? input) nil (data/str->int input)))

(defn int-match [re s]
  (parse-int (second (re-matches re s))))

(defn parse-naps [input]
  (loop [lines     (string/split-lines input)
         guard-id  -1
         asleep-at -1
         result    []]
    (if (empty? lines)
      result
      (let [line   (first lines)
            lines  (rest lines)
            guard  (int-match #"\[.*\] Guard #(\d+) .*" line)
            asleep (int-match #"\[.*:(\d+)\] falls asleep" line)
            awake  (int-match #"\[.*:(\d+)\] wakes up" line)]
        (cond (some? guard)
              (recur lines guard -1 result)

              (some? asleep)
              (recur lines guard-id asleep result)

              (some? awake)
              (recur lines guard-id -1
                     (conj result {guard-id (range asleep-at awake)})))))))

(defn merge-naps [parsed]
  (apply (partial merge-with concat) parsed))

(defn analyze-sleep [[id nap-minutes]]
  (let [stats          (frequencies nap-minutes)
        max-val        (apply max (vals stats))
        sleepy         (first (filter #(= max-val (second %)) stats))
        sleepy-minute  (first sleepy)
        naps-on-minute (second sleepy)]
    {:checksum                 (* id sleepy-minute)
     :total-minutes-slept      (count nap-minutes)
     :naps-on-sleepiest-minute naps-on-minute}))

(defn sleepiest-guard [sort-fn input]
  (->> (parse-naps input)
       merge-naps
       (map analyze-sleep)
       (sort-by sort-fn)
       last))

(defn part1 [input]
  (:checksum (sleepiest-guard :total-minutes-slept input)))

(defn part2 [input]
  (:checksum (sleepiest-guard :naps-on-sleepiest-minute input)))
