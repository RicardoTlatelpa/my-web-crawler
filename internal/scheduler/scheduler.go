package scheduler

import (
	"container/heap"
	"strings"
	"sync"
)


type URLItem struct {
	URL string
	Priority int
	index int
}

type PriorityQueue []*URLItem

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i,j int) bool {
	return pq[i].Priority > pq[j].Priority
}

func (pq PriorityQueue) Swap(i,j int) {
	pq[i],pq[j] = pq[j],pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*URLItem)
	item.index = len(*pq)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}

type Scheduler struct {
	seen map[string]bool
	pq PriorityQueue
	mu sync.Mutex
}

func New() * Scheduler{
	return &Scheduler{
		seen: make(map[string]bool),
		pq: make(PriorityQueue,0),
	}
}

func Score(url string) int {
	switch {
	case strings.Contains(url, "/scripts/") && strings.HasSuffix(url, ".html"):
		return 100
	case strings.Contains(url, "/genre"):
		return 90
	case strings.Contains(url, "all scripts"):
		return 80
	case strings.Contains(url, "imsdb.com"):
		return 50
	default:
		return 10
	}
}

func (s *Scheduler) ShouldVisit(url string) bool {
	s.mu.Lock()
	defer s.mu.Unlock()

	_, exists := s.seen[url]
	return !exists
}

func (s *Scheduler) MarkVisited(url string){
	s.mu.Lock()
	defer s.mu.Unlock()

	s.seen[url] = true
}

func (s *Scheduler) Enqueue(url string, priority int) {
	s.mu.Lock()
	defer s.mu.Unlock()
	if s.seen[url] {
		return
	}

	s.seen[url] = true
	heap.Push(&s.pq, &URLItem{
		URL: url,
		Priority: priority,
	})
}

func (s *Scheduler) Dequeue() string {
	s.mu.Lock()
	defer s.mu.Unlock()

	if s.pq.Len() == 0 {
		return ""
	}

	item := heap.Pop(&s.pq).(*URLItem)
	return item.URL
}