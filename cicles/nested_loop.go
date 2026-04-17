package main

import "fmt"

func main() {

outloop: //назвали внешний цикл outloop
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 2; j++ {
			fmt.Println(i, j)
			//break - прервет именно тот циклл который он лежит
			break outloop //внешний цикл остановится
			//так работает с continue
		}
	}

}

func gotoEx() {
	fmt.Println(1)
	goto HERE //переместись на строку 23 (там где стоит HERE)
HERE:
	fmt.Println(3)
}
