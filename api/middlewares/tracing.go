package middlewares

import (
	"context"
	"sync"
)

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

type traceIdKey struct{}

func SetTraceId(ctx context.Context, traceId int) context.Context {
	return context.WithValue(ctx, traceIdKey{}, traceId)
}

func GetTraceId(ctx context.Context) int {
	id := ctx.Value(traceIdKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}

	return 0
}
