package limiter

type Limiter interface {
	Allow() bool
}

type Task struct {
	dummy string
}
