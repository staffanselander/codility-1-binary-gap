package main

func Bits() (nums []int) {
	const MaxBits = 32

	for {
		if len(nums) == MaxBits {
			break
		}

		if len(nums) == 0 {
			nums = append(nums, 1)
			continue
		}

		next := nums[0] * 2
		nums = append([]int{next}, nums...)
	}

	return
}

func SumInts(nums []int) (sum int) {
	for _, value := range nums {
		sum += value
	}
	return
}
