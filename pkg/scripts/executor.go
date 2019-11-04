package scripts

import "sync"

var (
	r    *executor
	once sync.Once
)

type executor struct {
	mu sync.RWMutex
}

//Executor gives access to singleton executor
func Executor() *executor {
	once.Do(func() {
		r = &executor{}
	})

	return r
}
