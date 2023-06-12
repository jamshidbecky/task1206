package main

import (
	"fmt"
	"vazifas/if"
	"vazifas/func"
)

func main() {
	fmt.Println(ifMain.KabisaYili(100))
	fmt.Println(ifMain.Arifmetika(1, 2))
	ifMain.ErtangiSana(30, 12)
	ifMain.TubYokiTubEmas(21)

	fmt.Println("*********************")


	fmt.Println(funcMain.Ishora(1) + funcMain.Ishora(0))
	fmt.Println(funcMain.KvadraTenglamaIldizlari(1, -1, -6))
}