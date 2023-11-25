package limiter

type Limiter interface {
	TryEnqueue(task Task) EnqueueResult
}

type EnqueueResult struct {
	Enqueued bool
}

type Task struct {
	dummy string
}
