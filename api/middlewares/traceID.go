package middlewares

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

func newTraceID() int {
	return <-ids
}
