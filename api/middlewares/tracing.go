package middlewares

import "sync"

var (
	logNo int = 1
	mutex sync.Mutex
)

func newTraceId() int {
	var no int

	mutex.Lock()
	no = logNo
	logNo += 1
	mutex.Unlock()

	return no
}
