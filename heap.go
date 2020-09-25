package main

type IDDist struct {
	ID   string
	Dist float64
}

type MaxHeap struct {
	data []IDDist
}

func (h MaxHeap) Len() {
	return len(h.data)
}

func (h MaxHeap) Swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func (h MaxHeap) Less(i, j int) bool {
	return h.data[i].Dist > h.data[j].Dist
}

func (h *MaxHeap) Pop() interface{} {
	res := h.data[len(h.data)-1]
	h.data = h.data[:len(h.data)-1]
	return res
}

func (h *MaxHeap) Push(v interface{}) {
	h.data = append(h.data, v.(IDDist))
}
