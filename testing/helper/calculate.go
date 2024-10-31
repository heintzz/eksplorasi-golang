package helper

func Sum(nums ...int) (total int) {
	for _, n := range nums {
		total += n
	}

	return
}
