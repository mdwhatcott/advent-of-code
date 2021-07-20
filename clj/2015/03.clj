(ns advent.day03)

(require '[clojure.set :refer [union]])

(def directions {
  \< '(-1  0)
  \> '( 1  0)
  \^ '( 0  1)
  \v '( 0 -1)
})

(defn move [from direction]
  (map + from (directions direction)))

(defn unique-visits [arrows]
  (loop [at '(0 0)
         visits #{'(0 0)}
         steps (seq arrows)]
    (if (empty? steps)
        visits
        (let [next (move at (first steps))]
          (recur next
                 (conj visits next)
                 (rest steps))))))

(defn tag-team-unique-visits [arrows]
  (let [santa (take-nth 2 arrows)
        robot (take-nth 2 (rest arrows))]
    (union (unique-visits santa)
           (unique-visits robot))))

(defn part1 [arrows]
  (count (unique-visits arrows)))

(defn part2 [arrows]
  (count (tag-team-unique-visits arrows)))

(require '[clojure.test :refer :all])

(deftest day3

  ;; starting house and house to the right
  (testing "example-1" (is (= 2 (part1 ">"))))

  ;; square of 4 houses
  (testing "example-2" (is (= 4 (part1 "^>v<"))))

  ;; alternating two houses, up then down
  (testing "example-3" (is (= 2 (part1 "^v^v^v^v^v"))))

  (let [steps (slurp "03.txt")]
    (testing "part 1" (is (= 2572 (part1 steps))))
    (testing "part 2" (is (= 2631 (part2 steps))))))

(run-tests)