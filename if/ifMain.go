package ifMain

func KabisaYili(num int) string{
	flag := ""
	if num % 4 == 0 {
		if num % 400 == 0 {
			flag = "Bu yil kabisa yili"
		} else if num % 100 == 0 {
			flag = "Bu yil kabasi yili emas"
		} else {
			flag = "Bu yil kabisa yili"
		}
	} else {
		flag = "Bu yil kabasi yili emas"
	}
	return flag
}

func Arifmetika(x, y float64) (num1, num2 float64) {

	if x == y {
		num1 = x
		num2 = y
	} else if x < y {
		num1 = (x +y) / 2
		num2 = 2 * x * y 
	}

	
	return num1, num2
}