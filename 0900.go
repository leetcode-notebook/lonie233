package main

//type RLEIterator struct {
//	index int
//	Vals  []int
//}
//
//func ConstructorRLEIterator(A []int) RLEIterator {
//	var iterator = RLEIterator{0, nil}
//	var iteratorNumbers []int
//	for i := 0; i < len(A); i += 2 {
//		for j := 0; j < A[i]; j++ {
//			iteratorNumbers = append(iteratorNumbers, A[i+1])
//		}
//	}
//
//	iterator.Vals = iteratorNumbers
//	return iterator
//}
//
//func (this *RLEIterator) Next(n int) int {
//
//	if n+this.index > len(this.Vals)-1 {
//		return -1
//	}
//
//	this.index += n
//	return this.Vals[this.index-1]
//
//}

type RLEIterator struct {
	//当前计数的index
	curIndex int
	//当前index 已经减掉的值
	sub     int
	numbers []int
}

func ConstructorRLEIterator(A []int) RLEIterator {
	return RLEIterator{0, 0, A}
}

func (this *RLEIterator) Next(n int) int {

	if this.curIndex > len(this.numbers)-1 {
		return -1
	}

	if n+this.sub <= this.numbers[this.curIndex] {
		this.sub = this.sub + n
		return this.numbers[this.curIndex+1]
	}

	for n+this.sub > this.numbers[this.curIndex] {

		if this.curIndex > len(this.numbers)-1 {
			return -1
		}

		this.sub = (n + this.sub) - this.numbers[this.curIndex]
		n = 0
		this.curIndex = this.curIndex + 2

		if this.curIndex > len(this.numbers)-1 {
			return -1
		}

	}

	return this.numbers[this.curIndex+1]

}
