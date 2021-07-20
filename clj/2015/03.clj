(ns advent.day02)

(def directions {
  \< '(-1 0)
  \> '(1 0)
  \^ '(0 1)
  \v '(0 -1)
})

(defn move [from direction]
  (map + from (directions direction)))

(defn unique-visits [at visits steps]
  (if (empty? steps)
      (count visits)
      (let [next (move at (first steps))]
        (unique-visits next (conj visits next) (rest steps)))))

(defn count-unique-visits [steps]
  (unique-visits '(0 0) #{'(0 0)} steps))


(require '[clojure.test :refer :all])

(deftest day2
  (testing "example-1" (is (= 2 (count-unique-visits (seq ">")))))
  (testing "example-2" (is (= 4 (count-unique-visits (seq "^>v<")))))
  (testing "example-3" (is (= 2 (count-unique-visits (seq "^v^v^v^v^v")))))

  (let [steps (seq (slurp "03.txt"))]
    (testing "part 1" (is (= 2572 (count-unique-visits steps))))
  ))

(run-tests)