package funcMain

import "math"

func Ishora(num3 float64) (res int) {
	if num3 == 0 {
		res = 0
	} else if num3 > 0 {
		res = 1
	} else {
		res = -1
	}

	return res
}

func KvadraTenglamaIldizlari(A, B, C float64) (string) {
	flag := ""
	D := B * B - 4 * A * C
	if A == 0 || D == 0 {
		flag = "Tenglama bitta yechimga ega."
	} else if D > 0 {
		flag = "Tenglama ikkita yechimga ega."
	} else {
		flag = "Tenglama yechimga ega emas."
	}
	
	return flag
}

func DoiraYuziHisobla(r float64) (S float64) {
	S = r * r * math.Pi
	return S
}

func RingS(r1, r2 float64) (sFarq float64) {
	sFarq = (r1 * r1- r2 * r2) * math.Pi
	return sFarq
}

func TriangleP(a, b float64) float64 {
	c := math.Sqrt(a * a + b * b)

	perimetr := a + b + c

	return perimetr
}

func SumRange(n1, n2 int) (summ  int) {
	summ = 0
	if n1 > n2 {
		summ = 0
	} else {
		for n1 <= n2 {
			summ += n1
			n1 += 1
		}
	}
	return summ
}

func Calc(n3, n4, ishora int) (result int) {
	
	switch ishora {
	case 1: 
		result = n3 - n4
	case 2:
		result = n3 * n4
	case 3:
		result = n3 / n4
	default :
		result = n3 + n4
	}
	
	return result
}

