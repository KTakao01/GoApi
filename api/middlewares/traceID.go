package middlewares

import "context"

var (
	ids = make(chan int)
)

func init() {
	go func() {
		for i := 1; ; i++ {
			ids <- i
		}
	}()
}

type traceIDKey struct{}

func newTraceID() int {
	return <-ids
}

func SetTraceID(ctx context.Context, traceID int) context.Context {
	return context.WithValue(ctx, traceIDKey{}, traceID)
}

func GetTraceID(ctx context.Context) int {
	id := ctx.Value(traceIDKey{})

	if idInt, ok := id.(int); ok {
		return idInt
	}
	return 0
}
