#!/usr/bin/env bb

(require '[clojure.string :as str])

(def sample-a [
  "RL"
  ""
  "AAA = (BBB, CCC)"
  "BBB = (DDD, EEE)"
  "CCC = (ZZZ, GGG)"
  "DDD = (DDD, DDD)"
  "EEE = (EEE, EEE)"
  "GGG = (GGG, GGG)"
  "ZZZ = (ZZZ, ZZZ)"]) ; part 1: 2 steps

(def sample-b [
  "LLR"
  ""
  "AAA = (BBB, BBB)"
  "BBB = (AAA, ZZZ)"
  "ZZZ = (ZZZ, ZZZ)"]) ; part 1: 6 steps

(def sample-c [
  "LR"
  ""
  "11A = (11B, XXX)"
  "11B = (XXX, 11Z)"
  "11Z = (11B, XXX)"
  "22A = (22B, XXX)"
  "22B = (22C, 22C)"
  "22C = (22Z, 22Z)"
  "22Z = (22B, 22B)"
  "XXX = (XXX, XXX)"]) ; part 2: 6 steps

(def input-lines (str/split-lines (slurp "08.txt")))

(defn parse-terrain-line [line]
  (let [matches (re-seq #"[A-Z,0-9]{3}" line)]
    [(first matches) (rest matches)]))

(defn parse-terrain [lines]
  (into {} (map parse-terrain-line lines)))

(defn traverse [terrain directions start]
  (loop [at start path [] directions directions]
    (if (str/ends-with? at "Z")
        path
        (let [l-r   (first directions)
              which (if (= l-r \L) first second)
              next  (which (get terrain at))]
          (recur next (conj path next) (rest directions))))))

(defn count-steps [terrain directions from]
  (count (traverse terrain (cycle directions) from)))

(defn starting-nodes [terrain]
  (->> terrain keys (filter #(str/ends-with? % "A")) set))

(defn gcd [a b] (if (zero? b) a (recur b (mod a b))))
(defn lcm [a b] (/ (* a b) (gcd a b)))

(defn count-steps-simul [lines]
  (let [directions (first lines)
        terrain    (parse-terrain (drop 2 lines))
        starts     (starting-nodes terrain)
        counts     (map (partial count-steps terrain directions) starts)]
    (reduce lcm counts)))

(require '[clojure.test :refer :all])

(def sample-a-terrain (parse-terrain (drop 2 sample-a)))
(def sample-b-terrain (parse-terrain (drop 2 sample-b)))
(def sample-c-terrain (parse-terrain (drop 2 sample-c)))
(def input-terrain    (parse-terrain (drop 2 input-lines)))

(deftest tests
  (is (= ["AAA" ["BBB" "CCC"]] 
         (parse-terrain-line "AAA = (BBB, CCC)")))
  
  (is (= (hash-map "AAA" ["BBB" "CCC"]
                   "BBB" ["DDD" "EEE"]
                   "CCC" ["ZZZ" "GGG"]
                   "DDD" ["DDD" "DDD"]
                   "EEE" ["EEE" "EEE"]
                   "GGG" ["GGG" "GGG"]
                   "ZZZ" ["ZZZ" "ZZZ"]) 
         (parse-terrain (drop 2 sample-a))))
  
  (is (= ["CCC" "ZZZ"] 
         (traverse sample-a-terrain (cycle (first sample-a)) "AAA")))

  ;            L     L     R     L     L     R   
  (is (= ["BBB" "AAA" "BBB" "AAA" "BBB" "ZZZ"]
         (traverse sample-b-terrain (cycle (first sample-b)) "AAA")))

  (is (= 2 (count-steps sample-a-terrain (first sample-a) "AAA")))
  (is (= 6 (count-steps sample-b-terrain (first sample-b) "AAA")))
  (is (= 23147 (count-steps input-terrain (first input-lines) "AAA")))

  (is (= (set ["11A" "22A"]) (starting-nodes sample-c-terrain)))

  (is (= 12 (reduce lcm [2 3 4 6])))
  (is (= 6 (count-steps-simul sample-c)))
  (is (= 22289513667691 (count-steps-simul input-lines))))

(let [{:keys [fail error]} (run-tests)]
  (if (pos? (+ fail error)) (System/exit 1) (println "\nOK")))