package exercices

import (
	"fmt"
)

func Exercices() {
	exerciceOne()
	exercicesTwo()
	exerciceThree()
	exerciceFour()
	exerciceFive()
	exerciceSix()
	exerciceSeven()
	exerciceEight()
	exerciceNine()
	exerciceTen()
	exerciceEleven()
	exerciceTwelve()
	exerciceThirteen()
	exerciceFourteen()
	exerciceFifteen()
}

func exerciceOne() {
	fmt.Println("\nexercice One")
	x := 42
	y := " James Bond "
	z := true

	fmt.Print(x, y, z, "\n")
}

var x int
var y string
var z bool

func exercicesTwo() {
	fmt.Println("\nexercice Two")
	x = 42
	y = "James Bond"
	z = true

	fmt.Printf("%d %s %t \n", x, y, z)
}

var x2 int
var y2 string
var z2 bool

func exerciceThree() {
	fmt.Println("\nexercice Three")
	x2 = 42
	y2 = " James Bond "
	z2 = true

	s := fmt.Sprint(x2, y2, z2)
	fmt.Println(s)
}

type four int

var x3 four

func exerciceFour() {
	fmt.Println("\nexercice Four")
	x3 = 33
	fmt.Printf("%d\n", x3)
	fmt.Printf("%T\n", x3)
	x3 = 42
	fmt.Printf("%d\n", x3)
}

var y5 int

func exerciceFive() {
	fmt.Println("\nexercice Five")
	y5 = 33
	fmt.Printf("%d\n", y5)
	fmt.Printf("%T\n", y5)
	y5 = 42
	fmt.Printf("%d\n", y5)
	x3 = four(y5)
	fmt.Printf("%d\n", x3)
	fmt.Printf("%T\n", x3)
}

func exerciceSix() {
	fmt.Println("\nexercice Six")

	x6 := 7

	//mostra valroes em decimal, binario e hexa
	fmt.Printf("%d %#x %b\n", x6, x6, x6)
}

func exerciceSeven() {
	fmt.Println("\nexercice Seven")

	a7 := 4 == 6
	b7 := 4 > 3
	c7 := 6 >= 5
	d7 := 8 < 5
	e7 := 9 <= 11

	fmt.Println(a7, b7, c7, d7, e7)
}

func exerciceEight() {
	fmt.Println("\nexercice Eight")

	const x8 int = 10

	const y8 = "olá"

	fmt.Println(x8, y8)

}

func exerciceNine() {
	fmt.Println("\nexercice Nine")
	x9 := 10
	fmt.Printf("%d %#x %b\n", x9, x9, x9)
	y9 := x9 << 1
	fmt.Printf("%d %#x %b\n", y9, y9, y9)

	//percebe-se que é possivel dividir ou multiplar pela metade ou pelo dobro respectivamente ao deslocar 1 bit
}

func exerciceTen() {
	fmt.Println("\nexercice Ten")
	//raw string
	x10 := `olá! \n Tudo bem? \t -Sim!`
	//interpreted string
	y10 := "Olá!\nTudo bem? \t -Sim!"
	fmt.Println(x10)
	fmt.Println(y10)
}

func exerciceEleven() {
	fmt.Println("\nexercice Eleven")

	const (
		x11 = iota + 2023
		y11
		z11
	)

	fmt.Println(x11, y11, z11)
}

func exerciceTwelve() {
	fmt.Println("\nexercice Twelve")

	for x := 10; x <= 15; x++ {
		y := x % 4
		fmt.Print(y, " ")
	}
	fmt.Print("\n")
}

var arr [5]int

func exerciceThirteen() {
	fmt.Println("\nexercice Thirteen")

	for x := 0; x < 5; x++ {
		arr[x] = x
	}

	for _, value := range arr {
		fmt.Print(value, ", ")
	}

	fmt.Print("\n")
}

func exerciceFourteen() {
	fmt.Println("\nexercice Fourteen")

	x := 10 //armazena o valor 10
	y := &x //armazena o endereço de memoria de x
	z := *y //aponta o valor 10 contido no endereço de memoria contido em y

	a := &z //armazena o endereço de memoria de z
	b := *a //aponta o valor 10 que está sendo apontado em z que aponta para x

	b++

	fmt.Printf("O valor contido em x é %d e o valor contido em b é %d\n", x, b)
}

func exerciceFifteen() {
	fmt.Println("\nexercice Fifteen")

	x := 10
	y := 10

	//altera o valor de x sem retornar nada, apenas com ponteiro
	func(a *int) {
		*a += 5
	}(&x)

	//precisa retornar o valor da soma para atualizar y
	y = func(a int) int {
		a += 5
		return a
	}(y)

	fmt.Printf("O valor de x depois da função é %d e o valor de y depois da função é %d\n", x, y)
}
