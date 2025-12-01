package utils

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Sum(nums []int) int {
	total := 0
	for _, n := range nums {
		total += n
	}
	return total
}

func Product(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	result := 1
	for _, n := range nums {
		result *= n
	}
	return result
}

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func LCM(a, b int) int {
	return (a * b) / GCD(a, b)
}
func MinMax(nums []int) (min, max int) {
	if len(nums) == 0 {
		return 0, 0
	}
	min, max = nums[0], nums[0]
	for _, n := range nums[1:] {
		if n < min {
			min = n
		}
		if n > max {
			max = n
		}
	}
	return min, max
}

func Mod(a, b int) int {
    return ((a % b) + b) % b
}

func Divmod(numerator, denominator int) (quotient, remainder int) {
    quotient = numerator / denominator
    remainder = numerator % denominator
    return
}
