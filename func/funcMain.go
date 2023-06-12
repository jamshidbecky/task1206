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