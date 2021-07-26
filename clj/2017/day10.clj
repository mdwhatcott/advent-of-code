(ns day10)

(defn make-ring []
 (vec (take 256 (range))))

(defn twist [ring at length]
;;   4--5   pinch   4  5           4   1
;;  /    \  5,0,1  / \/ \  twist  / \ / \
;; 3      0  -->  3      0  -->  3   X   0
;;  \    /         \ /\ /         \ / \ /
;;   2--1           2  1           2   5
 (let [rotation  (take (count ring) (drop at (cycle ring)))
       slice     (take length rotation)
       reversed  (reverse slice)
       remaining (drop (count slice) rotation)
       complete  (concat reversed remaining)
       final     (take (count ring) (drop (- (count ring) at) (cycle complete)))]
  final))

(defn braid [original-ring lengths rounds]
 (loop [ring original-ring
        r    0
        at   0
        skip 0]
  (if (>= r (count lengths))
      ring
      (let [length    (nth lengths r)
            next-ring (twist ring at length)
            next-r    (inc r)
            next-at   (mod (+ at length skip) (count ring))
            next-skip (inc skip)]
       (recur next-ring next-r next-at next-skip)))))

(defn bi-product [ring]
 (* (first ring)
    (second ring)))


(require '[clojure.test :refer :all])

(deftest day1
 (testing "components"
  (is (= [2 1 0 3 4] (twist [0 1 2 3 4] 0 3)))
  (is (= [4 3 0 1 2] (twist [2 1 0 3 4] 3 4)))
  (is (= [4 3 0 1 2] (twist [4 3 0 1 2] 3 1)))
  (is (= [3 4 2 1 0] (twist [4 3 0 1 2] 1 5)))
  (is (= [3 4 2 1 0] (braid [0 1 2 3 4] [3 4 1 5] 1))))

 (testing "part 1"
  (let [part1-lengths [31 2 85 1 80 109 35 63 98 255 0 13 105 254 128 33]
        hashed-ring   (braid (make-ring) part1-lengths 1)]
   (is (= 6952 (bi-product hashed-ring)))))

 ;; https://www.lvguowei.me/post/nested-for-loops-in-clojure/
 (testing "part 2"
  (let [part2-lengths (seq (.getBytes (slurp "day10.txt")))
        extra-lengths '(17 31 73 47 23)
        all-lengths   (concat part2-lengths extra-lengths)
        hashed-ring   (braid (make-ring) all-lengths 64)]
   (is (= 42 42)))))

(run-tests)