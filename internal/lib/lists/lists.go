package lists

func Count(n int, l *[]int) int {
	count := 0
	for _, v := range *l {
		if v == n {
			count++
		}
	}
	return count
}
