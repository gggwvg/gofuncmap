package gofuncmap

import (
	"sync"
)

func init() {
	Add(stringFuncs)
}

var (
	// Funcs template function
	Funcs = map[string]interface{}{}

	// lock for adding user-defined functions
	mu sync.Mutex
)

// Add adds to values the functions in funcs. It does no checking of the input
func Add(funcMap map[string]interface{}) {
	mu.Lock()
	for name, fn := range funcMap {
		Funcs[name] = fn
	}
	mu.Unlock()
}
