(ns aoc.perf-spec
  (:require [speclj.core :refer :all]
            [aoc.perf :as sut])
  (:import (java.io StringWriter)))

(defmacro squelch-out
          "Like with-out-str, evaluates exprs in a context in which *out*
          is bound to a fresh StringWriter. Unlike with-out-str, returns
          the result of invoking the provided exprs and discards the string."
  [& body]
  `(let [s# (new StringWriter)]
     (binding [*out* s#]
       ~@body)))

(describe "Benchmarks"
  (with-stubs)

  (it "throws if x or n not positive"
    (should-throw IllegalArgumentException (sut/benchmark 0 1 "nope" #()))
    (should-throw IllegalArgumentException (sut/benchmark 1 0 "nope" #()))
    (should-throw IllegalArgumentException (sut/benchmark 0 0 "nope" #()))
    (should-throw IllegalArgumentException (sut/benchmark -1 1 "nope" #()))
    (should-throw IllegalArgumentException (sut/benchmark 1 -1 "nope" #()))
    (should-throw IllegalArgumentException (sut/benchmark -1 -1 "nope" #())))

  (it "invokes the provided fn and returns the average execution time"
    (let [x       10
          n       20
          fn      (stub :fn)
          average (squelch-out (sut/benchmark x n "Nop" fn))]
      (should (> average 0))
      (should-have-invoked :fn {:with [] :times (* x n)})))

  (it "emits a report detailing what was benchmarked"
    (with-redefs [sut/time-it (stub :time-it {:return 100})]
      (let [report (with-out-str (sut/benchmark 10 20 "Title" #()))]
        (should= "- Title - 10 x 20 ops, average per-op: 5ns\n" report)))))

