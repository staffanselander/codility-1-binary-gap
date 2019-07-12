package main

const MaxBits = 32

func Bits() (nums []int) {
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

func AsBinary(num int) (binary string) {
	for _, bitValue := range Bits() {
		if num >= bitValue {
			num -= bitValue
			binary += "1"
		} else if len(binary) > 0 {
			binary += "0"
		}
	}

	return
}
