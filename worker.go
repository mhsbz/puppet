package puppet

type Worker struct {
	pool *Pool
	task chan func()
}

func (w *Worker) run() {
	go func() {
		for f := range w.task {
			if f == nil {
				return
			}
			f()
		}
	}()
}
