package limiter

type Limiter interface {
	Allow(task Task) bool
}

type Task struct {
	dummy string
}
