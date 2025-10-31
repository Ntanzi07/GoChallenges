package data

import "sync"

type FibResult struct {
	Input    int
	Result   int
	Duration string
}

var (
	mu      sync.Mutex
	Results = make(map[int]FibResult)
)

func SaveResult(r FibResult) {
	mu.Lock()
	defer mu.Unlock()
	Results[r.Input] = r
}

func GetResult(n int) (FibResult, bool) {
	mu.Lock()
	defer mu.Unlock()

	r, ok := Results[n]
	return r, ok
}

func GetAll() []FibResult {
	mu.Lock()
	defer mu.Unlock()

	all := make([]FibResult, 0, len(Results))
	for _, v := range Results {
		all = append(all, v)
	}
	return all
}
