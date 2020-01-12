package main

type NumArray struct {
	sums []int
}

func Constructor(nums []int) NumArray {
	numArray := NumArray{}
	if len(nums) == 0 {
		return numArray
	}
	numArray.sums = make([]int, len(nums))
	numArray.sums[0] = nums[0]
	for i := 1; i < len(nums); i++ {
		numArray.sums[i] = numArray.sums[i-1] + nums[i]
	}
	return numArray
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sums[j]
	}
	return this.sums[j] - this.sums[i-1]
}
