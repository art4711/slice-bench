# Slicing a bench.

Just trying to figure out how much bounds checking Go can elide and if append is as efficient as doing the same thing manually. I'm mildly disappointed.

    BenchmarkInd-4      	       1	1122153289 ns/op	 267.34 MB/s
    BenchmarkRInd-4     	       3	 381926674 ns/op	 785.49 MB/s
    BenchmarkApp-4      	       2	 692376396 ns/op	 433.29 MB/s
    BenchmarkManual-4   	       3	 383093867 ns/op	 783.10 MB/s
    BenchmarkArr-4      	       3	 856799279 ns/op	 350.14 MB/s
    BenchmarkRArr-4     	       3	 494523622 ns/op	 606.64 MB/s
