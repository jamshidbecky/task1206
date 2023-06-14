package arrayMain

import "fmt"

func AppendEl(n, a, b int) (arr []int) {
	arr = []int{0, 0}
	arr[0] = a
	arr[1] = b
	for i := 2; i < n; i++ {
		var sum int
		for j := 0; j < i; j++ {
			sum += arr[j]
		}
		arr = append(arr, sum)
	}
	return
}

func ReverseArray(n int) ([]int) {
	var beforeArray []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		beforeArray = append(beforeArray, el)
	}

	var afterArray []int
	for i := n - 1; i >= 0; i -= 1 {
		afterArray = append(afterArray, beforeArray[i])
	}
	return afterArray
}

func OddArray(nuu int) []int {
	var arrNomi []int
	for i := 0; i < nuu; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		arrNomi = append(arrNomi, el)
	}
	var arrNomiToq []int
	for _, v := range arrNomi {
		if v % 2 != 0 {
			arrNomiToq = append(arrNomiToq, v)
		}
	}


	return arrNomiToq
}

func EvenArray(nuu int) []int {
	var arrNomi []int
	for i := 0; i < nuu; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		arrNomi = append(arrNomi, el)
	}
	var arrNomiJuft []int
	for _, v := range arrNomi {
		if v % 2 == 0 {
			arrNomiJuft = append(arrNomiJuft, v)
		}
	}


	return arrNomiJuft
}

