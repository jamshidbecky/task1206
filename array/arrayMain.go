package arrayMain

func AppendEl(n, a, b int) (arr []int) {
	sum := 0
	arr[0] = a
	arr[1] = b
	for i := 2; i < n; i++ {
		sum = arr[i] + arr[i + 1]
		arr = append(arr, sum)
		arr[0] = a
		arr[1] = b
		sum = 0
	}


	return
}