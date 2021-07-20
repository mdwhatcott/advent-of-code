(ns advent.day03)

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
        (count visits)
        (let [next (move at (first steps))]
          (recur next
                 (conj visits next)
                 (rest steps))))))


(require '[clojure.test :refer :all])

(deftest day3
  (testing "example-1" (is (= 2 (unique-visits ">"))))
  (testing "example-2" (is (= 4 (unique-visits "^>v<"))))
  (testing "example-3" (is (= 2 (unique-visits "^v^v^v^v^v"))))

  (let [steps (slurp "03.txt")]
    (testing "part 1" (is (= 2572 (unique-visits steps))))
  ))

(run-tests)