(ns aoc.y2018.d04-spec
  (:require [speclj.core :refer :all]
            [aoc.y2018.d04 :as sut]
            [aoc.data :as data]
            [clojure.string :as string]))

(def sample-input
  (str "[1518-11-01 00:00] Guard #10 begins shift\n"
       "[1518-11-01 00:05] falls asleep\n"
       "[1518-11-01 00:25] wakes up\n"
       "[1518-11-01 00:30] falls asleep\n"
       "[1518-11-01 00:55] wakes up\n"
       "[1518-11-01 23:58] Guard #99 begins shift\n"
       "[1518-11-02 00:40] falls asleep\n"
       "[1518-11-02 00:50] wakes up\n"
       "[1518-11-03 00:05] Guard #10 begins shift\n"
       "[1518-11-03 00:24] falls asleep\n"
       "[1518-11-03 00:29] wakes up\n"
       "[1518-11-04 00:02] Guard #99 begins shift\n"
       "[1518-11-04 00:36] falls asleep\n"
       "[1518-11-04 00:46] wakes up\n"
       "[1518-11-05 00:03] Guard #99 begins shift\n"
       "[1518-11-05 00:45] falls asleep\n"
       "[1518-11-05 00:55] wakes up"))

(def real-input (data/read 2018 4))

(describe "2018 Day 4"
  (context "Part 1"
    (it "parses input"
      (->> (sut/parse-naps sample-input)
           (should=
             [{10 (range 5, 25)}
              {10 (range 30 55)}
              {99 (range 40 50)}
              {10 (range 24 29)}
              {99 (range 36 46)}
              {99 (range 45 55)}])))

    (it "merges input"
      (->> (sut/parse-naps sample-input)
           (sut/merge-naps)
           (should=
             {10 (concat (range 5, 25)
                         (range 30 55)
                         (range 24 29))
              99 (concat (range 40 50)
                         (range 36 46)
                         (range 45 55))})))

    (it "analyzes a guard #10's sleep patterns"
      (as-> (sut/parse-naps sample-input) $
            (sut/merge-naps $)
            (get $ 10)
            (sut/analyze-sleep [10 $])
            (should=
              $ {:checksum                 (* 10 24)
                 :total-minutes-slept      50
                 :naps-on-sleepiest-minute 2})))

    (it "analyzes a guard #99's sleep patterns"
      (as-> (sut/parse-naps sample-input) $
            (sut/merge-naps $)
            (get $ 99)
            (sut/analyze-sleep [99 $])
            (should=
              $ {:checksum                 (* 99 45)
                 :total-minutes-slept      30
                 :naps-on-sleepiest-minute 3})))

    (it "solves part 1 with sample data"
      (should= (* 10 24) (sut/part1 sample-input)))

    (it "solves part 1 with real data"
      (should= 35184 (sut/part1 real-input)))
    )

  (context "Part 2"
    (it "solves part 2 with sample data"
      (should= (* 99 45) (sut/part2 sample-input)))

    (it "solves part 2 with real data"
      (should= 37886 (sut/part2 real-input)))
    )
  )
