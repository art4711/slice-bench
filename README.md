# Slicing a bench.

Just trying to figure out how much bounds checking Go can elide and if append is as efficient as doing the same thing manually. I'm mildly disappointed.

    BenchmarkInd-4      	       1	1121821884 ns/op	 267.42 MB/s
    BenchmarkRInd-4     	       3	 378838563 ns/op	 791.89 MB/s
    BenchmarkApp-4      	       2	 685736934 ns/op	 437.49 MB/s
    BenchmarkManual-4   	       3	 383881650 ns/op	 781.49 MB/s