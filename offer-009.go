package main

type CQueue struct {
	elements []int
}

func ConstructorCQueue() CQueue {
	return CQueue{}
}

func (this *CQueue) AppendTail(value int) {
	this.elements = append(this.elements, value)
}

func (this *CQueue) DeleteHead() int {
	if len(this.elements) == 0 {
		return -1
	}
	value := this.elements[0]
	this.elements = this.elements[1:]
	return value
}
