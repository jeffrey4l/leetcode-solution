package solution

import (
	"fmt"
	"time"
)

func timeit(f func(), name string, times int) {
	started := time.Now()
	for i := 0; i < times; i++ {
		f()
	}
	during := time.Since(started)
	fmt.Printf("%s takes: %v\n", name, during)
}
