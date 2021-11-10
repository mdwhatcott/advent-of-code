(defproject aoc "0.1.0-SNAPSHOT"
  :description "FIXME: write description"
  :url "http://example.com/FIXME"
  :license {:name "Eclipse Public License"
            :url  "http://www.eclipse.org/legal/epl-v10.html"}
  :main aoc.core
  :dependencies [[org.clojure/clojure "1.8.0"]
                 [org.clojars.mdwhatcott/benchmarks "0.3.1"]]
  :profiles {:dev {:dependencies [[speclj "3.3.2"]]}}
  :plugins [[speclj "3.3.2"]]
  :test-paths ["spec"]
  :aliases {"aoc" ["run" "-m" "aoc.tools.generator"] #_"ie. lein aoc 2021 1"}
  )
