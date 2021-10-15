(ns aoc.perf)

(defmacro time-ns [ex]
  `(let [start# (. System (nanoTime))]
     (do ~ex (- (. System (nanoTime)) start#))))

(defn benchmark
  ([title f] (benchmark 10 10000 title f))
  ([n title f] (benchmark 10 n title f))
  ([x n title f]
   (let [times   (for [_ (range x)] (time-ns (dotimes [_ n] (f))))
         per-ops (map #(/ % n) times)
         average (int (/ (apply + per-ops) (count per-ops)))]
     (println
       (format "## %s - %d x %d ops, average per-op: %dns" title x n average)))))
