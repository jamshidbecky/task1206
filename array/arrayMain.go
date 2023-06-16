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

func Array5(n int) ([]int, []int) {
	var asosiyArray []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		asosiyArray = append(asosiyArray, el)
	}

	var juftArray, toqArray []int
	for _, v := range asosiyArray {
		if v % 2 == 0 {
			juftArray = append(juftArray, v)
		} else {
			toqArray = append(toqArray, v)
		}
	}

	min := juftArray[0]
	max := toqArray[0]

	for i := 0; i < len(juftArray) - 1; i++ {
		if min > juftArray[i] {
			min = juftArray[i]
			juftArray = append(juftArray, min)
		} 
	}

	for  j:= 0; j < len(toqArray) - 1; j++ {
		if max < toqArray[j] {
			max = toqArray[j]
			toqArray = append(toqArray, max)
		} 
	}

	return toqArray, juftArray
}

func Array6(n, k int) ([]int) {
	var array []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		array = append(array, el)
	}

	var arrayK []int 
	for i := k; i <= n; i += k {
		varib := array[i]
		arrayK = append(arrayK, varib)
	}

	
	return arrayK
}

func Array8(n int) (int) {
	var array []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		array = append(array, el)
	}

	elFirst := 0

	for i := 0; i < len(array); i++ {
		if array[i] < array[len(array) - 1] {
			elFirst = array[i]
			break
		}
	}
		
	return elFirst
}

func Array9(n int) (int) {
	var array []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		array = append(array, el)
	}

	elFirst := 0

	for i := len(array) - 1; i >= 0; i -= 1 {
		if array[len(array) - 1] > array[i] {
			if array[i] > array[0] {
			elFirst = array[i]
			}
			break
		}	
	}
		
	return elFirst
}

func Array11(n, k, l int) (float64) {
	var array []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i+1)
		var el int
		fmt.Scan(&el)
		array = append(array, el)
	}

	sum := 0

	for j := k; j <= l ; j++ {
		sum += array[j]
	}

	summ := float64(sum) / float64(l - k + 1)

	return summ
}

func Array13(n, k, l int) (float64) {
	var array []int
	for i := 0; i < n; i++ {
		fmt.Printf("%v-elementni kiritng: ", i)
		var el int
		fmt.Scan(&el)
		array = append(array, el)
	}

	sum := 0

	for j := 0; j <= k ; j++ {
		sum += array[j]
	}

	for c := l; c < n ; c++ {
		sum += array[c]
	}

	summ := float64(sum) / float64(k + 1 + n - l)

	return summ
}