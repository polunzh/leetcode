package numberofrecentcalls

type RecentCounter struct {
	requests []int
}

func Constructor() RecentCounter {
	return RecentCounter{}
}

func (this *RecentCounter) Ping(t int) int {
	first := t - 3000
	firstIndex := 0
	for i := firstIndex; i < len(this.requests) && this.requests[i] < first; i++ {
		firstIndex++
	}

	this.requests = append(this.requests, t)
	this.requests = this.requests[firstIndex:]

	return len(this.requests)
}
