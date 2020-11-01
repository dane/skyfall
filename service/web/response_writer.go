package web

type Contexter interface {
	Context() context.Context
	SetContext(context.Context)
}

type responseWriter struct {
	http.ResponseWriter

	mu  *sync.RWMutex
	ctx context.Context
}

func (rw *responseWriter) Context() context.Context {
	rw.mu.RLock()
	defer rw.mu.RUnlock()
	return rw.ctx
}

func (rw *responseWriter) SetContext(ctx context.Context) {
	rw.mu.Lock()
	defer rw.mu.Unlock()
	rw.ctx = ctx
}
