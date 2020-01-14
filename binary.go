package main

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
