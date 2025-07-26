package frontier

import "sync"

type Frontier struct {
	visited map[string] bool
	queue []string
	mu sync.Mutex
}

func New() *Frontier {
	return &Frontier{
		visited: make(map[string]bool),
		queue: make([]string, 0),
	}
}

func (f *Frontier) Enqueue(url string) {
	f.mu.Lock()
	defer f.mu.Unlock()
	
}