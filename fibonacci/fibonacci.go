// This was my attempt to implement separate functions within a single function.
// I was also trying to make the code as unreadable as possible.
// BTW, I thought this would be more efficient than using functions.
// It was not.

package fibonacci

func Fibonacci(n int) int {
	var (
		nums    []int
		current int
	)
	i := -1
	need := n
	goto MAIN
CALCULATE:
	if i >= 2 {
		nums = append(nums, nums[i-1]+nums[i-2])
	}
	current = nums[i]
	goto LOOP
RETURN:
	return current
MAIN:
	current = 0
	nums = make([]int, 2, 10)
	nums[0] = 1
	nums[1] = 1
LOOP:
	if i < need {
		i++
		goto CALCULATE
	}
	goto RETURN
}
