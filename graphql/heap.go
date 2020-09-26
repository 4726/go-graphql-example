package graphql

type idDist struct {
	ID   string
	Dist float64
}

//maxHeap implements the heap.Interface interface
type maxHeap struct {
	data []idDist
}

func (h maxHeap) Len() int {
	return len(h.data)
}

func (h maxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h maxHeap) Less(i, j int) bool {
	return h.data[i].Dist > h.data[j].Dist
}

func (h *maxHeap) Pop() interface{} {
	res := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return res
}

func (h *maxHeap) Push(v interface{}) {
	h.data = append(h.data, v.(idDist))
}
