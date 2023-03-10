package lessons

import (
	"encoding/json"
	"fmt"
	"os"
)

func Lessons() {
	lessonOne()
	lessonTwo()
	lessonThree()
	lessonFour()
	lessonFive()
	lessonSix()
	lessonSeven()
	lessonEight()
	lessonNine()
	lessonTen()
	lessonEleven()
	lessonTwelve()
	lessonThirteen()
	lessonFourteen()
	lessonFifteen()
	lessonSixTeen()
	lessonSeventeen()
	lessonEighteen()
	lessonNineteen()
	lessonTwenty()
	lessonTwentyOne()
	lessonTwentyTwo()
	lessonTwentyThree()
	lessonTwentyFour()
	lessonTwentyFive()
	lessonTwentySix()
	lessonTwentySeven()
	lessonTwentyEight()
	lessonTwentyNine()
	lessonThirty()
	lessonThirtyOne()
	lessonThirtyTwo()
	lessonThirtyThree()
	lessonThirtyFour()
	lessonThirtyFive()
	lessonThirtySix()
}

// a nivel de package, apenas funciona atribuição. Gopher não funciona. O valor não precisa obrigatriamente ser usado
var myAge = 25

func lessonOne() {

	fmt.Println("Lesson One")

	//operador gopher serve para declarar variaveis com tipagem dinâmica
	// se tenta-se colocar var myName := "blablabla" daria erro pois o gopher se encarrega de tipar
	//variaveis tem contexto local (func)
	myName := "Mateus"

	fmt.Println("Meu nome", myName)

	//o gopher não pode ser utilizado p atribuir. deve-se usar =

	myName = "Mateus Vedoy"

	fmt.Println("Meu nome", myName)

	//cenário 1: quero pegar o retorno de algo para usar
	byteNumber, err := fmt.Println("Helo Word")
	fmt.Printf("Quantidade de byestes usados %d e quantidade de erros retornados %s", byteNumber, err)

	//cenário 2: não quero utilizar algo pois nao me interessa
	byteNumber2, _ := fmt.Println("Helo Word")
	fmt.Printf("Quantidade de byestes usados %d\n", byteNumber2)
}

// declarando variaveis com tipagem forte.
var age int = 25
var age2 = 25

// caso uma variavel seja apenas delacara em escopo de pacote, só podera ser atribuido valor dentro de um bloco (func)
var age3 int

func lessonTwo() {
	fmt.Println("\nLesson Two")

	age3 = 25

	//O tipo int é opcional, pois de qualquer forma o tipo sera o mesmo para esse valor inteiro
	fmt.Printf("%T, %T\n", age, age2)
}

// constantes
func lessonThree() {
	fmt.Println("\nLesson Three")
	//o tipo de uma constante é definido apenas quando usado. Diferente de uma variavel q qnd criada já é tipada
	const x = 23

	//pode declarar const assim tbm
	const (
		y = 2
		z = 4
	)

	fmt.Println(x, y, z)
}

// iota
func lessonFour() {

	fmt.Println("\nLesson Four")
	//serve para gerar numeros inteiros não tipados e sucessivos

	const (
		x = iota
		y = iota
		z = iota
	)

	fmt.Println(x, y, z)

	//tambem pode ser inferido apenas no primeiro elemento e com operações matematicas

	const (
		a = iota * 3
		b
		c
	)

	fmt.Println(a, b, c)
}

// laços de repetição
func lessonFive() {
	fmt.Println("\nLesson Five")

	for x5 := 0; x5 < 2; x5++ {
		fmt.Print(x5, " ")
	}
}

// usando for como while
func lessonSix() {
	fmt.Println("\nLesson Six")

	x6 := 0

	for x6 <= 3 {
		fmt.Print(x6, " ")
		x6++
	}
}

// if else if else
func lessonSeven() {
	fmt.Println("\nLesson Seven")

	x := 4

	//obrigatoriamente o else deve estar na mesma linha do }, pois do contrário não roda

	if x <= 4 {
		fmt.Println(x)
	} else if x > 4 {
		fmt.Printf("%#x", x)
	} else {
		fmt.Printf("%b", x)
	}
}

// switch
func lessonEight() {
	fmt.Println("\nLesson Eight")
	x := 9

	//break não é necessário
	// usar fallthrough para pular a validação do case abaixo do que cair verdadeiro
	switch x {
	case 3:
		fmt.Println("Três")
		fallthrough //se o valor de x fosse 3, mostraria três e nove (mas não validaraia o case 9)
	case 9:
		fmt.Println("Nove")
	default:
		fmt.Println("Nenhum")
	}
}

// pegando tipo de uma interface desconhecida
var l9 interface{}

// type of
func lessonNine() {
	fmt.Println("\nLesson Nine")

	l9 = 11.36

	fmt.Printf("%T\n", l9)

	switch l9.(type) {
	case int:
		fmt.Printf("%d é um inteiro\n", l9)
	case bool:
		fmt.Printf("%t é um booleano\n", l9)
	case string:
		fmt.Printf("%s é um string\n", l9)
	case float64:
		fmt.Printf("%2.f é um flutuante\n", l9)
	default:
		fmt.Println("Não sei")
	}
}

// array
var l10 []int //isso aqui impede que se atribua valores mais além a esse array
var l102 [2]int

func lessonTen() {
	fmt.Println("\nLesson Ten")

	l103 := [3]string{"1", "2", "3"}

	fmt.Println(l10)
	fmt.Println(l102)
	fmt.Println(l103)

	fmt.Println(len(l10))
	fmt.Println(len(l102))
	fmt.Println(len(l103))

	l102[0] = 0
	l102[1] = 1

	fmt.Println(l10)
	fmt.Println(l102)
}

// slice
func lessonEleven() {
	fmt.Println("\nLesson Eleven")

	slice := []int{1, 2} //é preciso iniciar com valores

	slice[0] = 1
	slice[1] = 2

	fmt.Println(slice)
	fmt.Println(len(slice))

	//a forma de adicionar mais valores
	slice = append(slice, 8)

	fmt.Println(slice)
	fmt.Println(len(slice))

	slice[0] = 9 //altera o valor de algum indice do slice/array
	fmt.Println(slice)
}

// range of slice
func lessonTwelve() {
	fmt.Println("\nLesson Twelve")

	slice := []string{"a", "b"}

	for index, value := range slice {
		fmt.Print(index, ":", value, " | ")
	}

	fmt.Print("\n")
}

// slice no slice
func lessonThirteen() {
	fmt.Println("\nLesson Thirteen")

	slice := []string{"Mateus", "Vedoy", "Goes"}

	fmt.Println(slice)

	mateusVedoy := slice[0:2] //selecionando parte do slice anterior

	fmt.Println(mateusVedoy)

	mateus := append(slice[:1], slice[3:]...) //remove valores de indices e indices indesejados

	fmt.Println(mateus)

	slice2 := []string{"Mateus"}

	slice3 := []string{"Vedoy", "Goes"}

	slice4 := append(slice2, slice3...) //tem q ser assim pois slice3 tem varios valores

	fmt.Println(slice4)
}

// make do slice
func lessonFourteen() {
	fmt.Println("\nLesson Fourteen")

	//serve para realizar o processo de append no slice de forma automática ao atingir o limite atual do slice

	//make([]tipo, qtd de valores, tamanho máximo)
	slice := make([]int, 5, 6)

	fmt.Println("Valoes do slice: ", slice, " | capacidade máxima: ", cap(slice))

	slice[0], slice[1], slice[2], slice[3], slice[4] = 1, 2, 3, 4, 5

	fmt.Println("Valoes do slice: ", slice, " | capacidade máxima: ", cap(slice))

	slice = append(slice, 7)
	slice = append(slice, 8)

	fmt.Println("Valoes do slice: ", slice, " | capacidade máxima: ", cap(slice))

	//ao atingir o limite do slice, o cap é dobrado
}

//slice multidimensional

func lessonFifteen() {
	fmt.Println("\nLesson Fifteen")

	multSlice := [][]int{
		[]int{1, 2, 3},
		[]int{4, 5, 6},
	}

	fmt.Println(multSlice)
}

// append no slice
func lessonSixTeen() {
	fmt.Println("\nLesson Sixteen")
	//append acaba alterando na memória o mesmo local do slice original caso o append não estrapole seu tamanho

	slice1 := []int{1, 2, 3, 4, 5}

	fmt.Println("Antes do append: ", slice1)

	slice2 := append(slice1[:2], slice1[4:]...)

	fmt.Println("Depois do append: ", slice1)
	fmt.Println(slice2)

	slice3 := []int{4, 5, 6, 7}

	fmt.Println("Antes do append: ", slice3)

	slice4 := append(slice3, 8, 9)

	fmt.Println("Depois do append: ", slice3)
	fmt.Println(slice4)
}

// map
func lessonSeventeen() {
	fmt.Println("\nLesson Seventeen")

	maps := map[string]int{
		"zero": 0,
		"um":   1,
		"dois": 2,
	}

	fmt.Println(maps)
	fmt.Println(maps["1"])

	maps["tres"] = 3

	fmt.Println(maps)

	//ok trará false, pois não existe a chave buscada (5)
	value, ok := maps["cinco"]

	fmt.Println(value, ok)
}

// range e dletando valores de map
func lessonEighteen() {
	fmt.Println("\nLesson Eighteen")

	maps := map[int]string{
		0: "zero",
		1: "um",
		2: "dois",
	}

	//range
	for key, value := range maps {
		fmt.Println("Chave: ", key, " e valor: ", value)
	}

	//deletar
	//delete(mapa, indice)
	delete(maps, 1)

	fmt.Println(maps)
}

//struct

type l19 struct {
	x string
	y int
	z bool
}

func lessonNineteen() {
	fmt.Println("\nLesson Nineteen")

	a := l19{
		x: "xis",
		y: 1,
		z: true,
	}

	b := l19{
		"chis", 2, false,
	}

	fmt.Println(a, b)
	fmt.Println(a.x, b.y)
}

// struct interno
type l20 struct {
	a string
	b bool
}

type l201 struct {
	x int
	y l20
}

func lessonTwenty() {
	fmt.Println("\nLesson Twenty")

	a1 := l201{
		x: 12,
		y: l20{
			a: "123",
			b: false,
		},
	}

	a2 := l20{
		a: "1212",
		b: true,
	}

	fmt.Println(a1, a2)
	fmt.Println(a1.y.a, a2.b)
}

//struct anônimo

func lessonTwentyOne() {
	fmt.Println("\nLesson Twenty One")

	x := struct {
		a string
		b int
	}{"1123", 4}

	fmt.Println(x)
}

//funções

func noArgsAndNoReturn() {
	fmt.Println("Função sem parâmetros e sem retorno")
}

func withArgsButDoesntReturn(msg string) {
	fmt.Printf("Mensagem do parâmetro: %s\n", msg)
}

func withArgsAndOneReturn(msg string) string {
	return fmt.Sprint("Mensagem de retorno com parâmetro: ", msg)
}

func withVariadicArgsAndTwoReturn(values ...int) (int, string) {
	sum := 0
	for _, val := range values {
		sum += val
	}
	msg := fmt.Sprint("Quantidade de termos: ", len(values))

	return sum, msg
}

func lessonTwentyTwo() {
	fmt.Println("\nLesson Twenty two")

	noArgsAndNoReturn()
	msgReturn := withArgsAndOneReturn("Mensagem teste")
	fmt.Println(msgReturn)
	withArgsButDoesntReturn("Mensagem teste 2")
	sum, str := withVariadicArgsAndTwoReturn(1, 2, 3, 5, 7)
	fmt.Println("Valor da soma: ", sum, " e ", str)
}

// defer - adiar / deixar para o fim
func lessonTwentyThree() {
	fmt.Println("\nLesson Twenty three")

	//defer funcion como pilha, onde o primerio defer a entrar é o último a sair
	defer fmt.Println("Executa por fim")
	defer fmt.Println("Executa em penúltimo")
	fmt.Println("Executa primeiro")
}

//métodos

type somenthing struct {
	somenthingInt    int
	somenthingString string
}

// (s Somenthing) informa que essa função sera um método de Something apenas
func (s somenthing) somenthingToDo() {
	someInt := s.somenthingInt
	someString := s.somenthingString
	fmt.Printf("Somenthing tem uma prop int %d e uma prop string %s\n", someInt, someString)
}

func lessonTwentyFour() {
	fmt.Println("\nLesson Twenty Four")

	some := somenthing{12, "Alguma palavra"}
	some.somenthingToDo()
}

//interface e polimorfismo

type colorBlue struct {
	color string
}

type colorYellow struct {
	color string
}

// metodos com polimorfismo
func (c colorBlue) whatColor() {
	fmt.Printf("Sou dar cor %s e sou uma cor fria\n", c.color)
}

func (c colorYellow) whatColor() {
	fmt.Printf("Sou dar cor %s e sou uma cor quente\n", c.color)
}

// interface
type whatColor interface {
	whatColor()
}

// func q implementa interface
func sayWhatColor(w whatColor) {
	w.whatColor()
}

func lessonTwentyFive() {
	fmt.Println("\nLesson Twenty Five")

	azul := colorBlue{"Azul"}
	amarelo := colorYellow{"Amarelo"}

	//forma simples de chamar os métodos das struct
	azul.whatColor()
	amarelo.whatColor()

	//chamando via interface
	sayWhatColor(azul)
	sayWhatColor(amarelo)
}

// funções anônimas | rodam 1x apenas
func lessonTwentySix() {
	fmt.Println("\nLesson Twenty Six")

	//sem param e retorno
	func() {
		fmt.Println("Função anônima")
	}()

	withArgAndOneReturn := func(msg string) string {
		return fmt.Sprint("Mensagem anônima: ", msg)
	}("Textinho anônimo")

	fmt.Println(withArgAndOneReturn)

	op, soma, tam := func(operation string, values ...int) (string, int, int) {
		sum := 0
		msg := fmt.Sprint("Operação realizada: ", operation)

		for _, val := range values {
			sum += val
		}

		return msg, sum, len(values)
	}("Soma", 2, 4, 5, 7)

	fmt.Printf("Operação: %s, Resultado da soma: %d e quantidade de termos somados: %d\n", op, soma, tam)
}

// passando func como expressão p uma variavel
func lessonTwentySeven() {
	fmt.Println("\nLesson Twenty Seven")

	shouldSumBunchOfValues := func(values ...int) int {
		sum := 0
		for _, value := range values {
			sum += value
		}

		return sum
	}

	shouldPrintResult := func(val int) {
		fmt.Println("Soma final: ", val)
	}

	finalSum := shouldSumBunchOfValues(1, 2, 3, 4, 8)

	shouldPrintResult(finalSum)
}

//função que retorna função

func calculateAmountOfValuesAndReturnFunctionThatResolveValuesLength(values ...int) func() int {
	sum := 0

	for _, val := range values {
		sum += val
	}

	fmt.Println("O resultado dessa soma é :", sum)

	return func() int {
		return len(values)
	}
}

//pode retornar qualquer tipo de função com ou sem parametro

func lessonTwentyEight() {
	fmt.Println("\nLesson Twenty Eight")

	amount := calculateAmountOfValuesAndReturnFunctionThatResolveValuesLength(1, 2, 3, 4, 5, 6)

	amountLen := amount()

	fmt.Println("Tamanho do array passado: ", amountLen)

	amountLenTwo := calculateAmountOfValuesAndReturnFunctionThatResolveValuesLength(1, 2, 3, 4, 5, 6)()

	fmt.Println("Funciona do mesmo jeito! Tamanho do array: ", amountLenTwo)

}

//callback (passar função como argumento)

// função de soma para ser um callback
func sum(values ...int) int {
	sum := 0
	for _, val := range values {
		sum += val
	}
	return sum
}

// função que tem soma apenas pares e recebe uma função como param
func onlySumPairNumbers(f func(values ...int) int, pair ...int) int {
	var slice []int
	for _, value := range pair {
		if value%2 == 0 {
			slice = append(slice, value)
		}
	}
	return f(slice...)
}

func lessonTwentyNine() {
	fmt.Println("\nLesson Twenty Nine")

	sumOfValues := onlySumPairNumbers(sum, []int{23, 44, 11, 2, 3, 7, 8, 9, 12, 34}...)
	fmt.Println("Soma de pares: ", sumOfValues)
}

//closure

func innerContextToClosure() func() int {
	x := 0

	//essa func pegará o x do contexto q está inserida
	return func() int {
		x++
		return x
	}
}

func lessonThirty() {
	fmt.Println("\nLesson Thirty")

	first := innerContextToClosure()
	//cada chamada terá um valor diferente, pois dentro do escopo de fisrt, x está sendo alterado
	fmt.Printf("1) %d; 2) %d; 3) %d\n", first(), first(), first())

	second := innerContextToClosure()
	//como se trata de outra variavel e por isso outro contexto, os valores reiniciam do zero novamente
	fmt.Printf("1) %d; 2) %d; 3) %d\n", second(), second(), second())
}

//recurisividade

func fat(value int) int {
	if value == 0 {
		return 1
	} else {

		return value * fat(value-1)
	}
}

func lessonThirtyOne() {
	fmt.Println("\nLesson Thirty One")

	fmt.Println("Fatorial de 5 é: ", fat(5))
}

// ponteiro
func lessonThirtyTwo() {
	fmt.Println("\nLesson Thirty Two")

	//cria variavel com valor
	x := 10

	//atribui o endereco de memoria de x ao y
	y := &x

	//atribui o valor do endereço de memória y que por sua vez contém o endereço de memoria de x
	//no fim z guardará o valor contido em x
	z := *y

	fmt.Printf("Variavel x tem o valor: %d e está armazenada no endereço de memória %d\n", z, y)
	//*int siginifica ponteiro de inteiro
	fmt.Printf("O tipo de x é %T, o tipo de y é %T e o tipo de z é %T\n", x, y, z)
}

//quando usar ponteiro
//muito útil para poupar performance ao usar variaveis ou copiar valores já definidos no programa

// aqui o valor de x será copiado. O que gerará um gasto extra de performance
func thisReceiveAValue(x int) {
	x++
	fmt.Printf("Valor de x dentro da função é: %d\n", x)
}

// aqui y aponta direto para o valor contido no endereço de y
func thisReceiveAPointer(y *int) {
	*y++
	fmt.Printf("O valor de y dentro da função é: %d\n", *y)
}

func lessonThirtyThree() {

	//qual a diferença afinal?
	// sempre que o valor "original" pecisar ser mantido, não use ponteiro, mas quando o valor
	//"original" puder mudar, use
	fmt.Println("\nLesson Thirty Three")

	x := 0
	y := 0

	thisReceiveAValue(x)
	thisReceiveAPointer(&y)

	fmt.Printf("O valor de x após a função é %d e o valor de y após a função é %d\n", x, y)
}

//transformar struct em json

// struct tem q ter letra maiuscula p ser exportada e visivel p outros pacotes
// o mesmo serve para as props
type Category struct {
	Name string
}

type Product struct {
	Name      string
	UnitValue float64
	Brand     Brand
	Category  Category
}

type Brand struct {
	Name string
}

func lessonThirtyFour() {
	fmt.Println("\nLesson Thirty Four")

	pasta := Product{
		Name:      "Spaghetti",
		UnitValue: 3.59,
		Brand:     Brand{"Liza"},
		Category: Category{
			"Food",
		},
	}

	fmt.Println("A struct criada foi ", pasta)

	byteJson, err := json.Marshal(pasta)

	if err != nil {
		fmt.Println("Algo deu errado: ", err)
	} else {
		fmt.Print("O json gerado foi: ")
		fmt.Print(string(byteJson))
		fmt.Print("\n")
	}
}

// JSON to struct
func lessonThirtyFive() {
	fmt.Println("\nLesson Thirty Five")

	sliceOfBytes := []byte(`{"Name":"Spaghetti","UnitValue":3.59,"Brand":{"Name":"Liza"},"Category":{"Name":"Food"}}`)

	type Product struct {
		Name      string  `json:"Name"` //`json:Name` é uma tag para dizer que a prop Name do json será Name tbm na struct
		UnitValue float64 `json:"UnitValue"`
		Brand     struct {
			Name string `json:"Name"`
		} `json:"Brand"`
		Category struct {
			Name string `json:"Name"`
		} `json:"Category"`
	}

	var pasta Product

	err := json.Unmarshal(sliceOfBytes, &pasta)

	if err != nil {
		fmt.Println("Erro: ", err)
	}
	fmt.Println(pasta)
}

// criando Enconder JSON
func lessonThirtySix() {
	fmt.Println("\nLesson Thirty Six")

	pasta := Product{
		Name:      "Spaghetti",
		UnitValue: 3.59,
		Brand:     Brand{"Liza"},
		Category: Category{
			"Food",
		},
	}

	fmt.Println("Struct pasta: ", pasta)

	//os.Stout já ira converter o valor da struct e tocar na tela (os.Stdout)
	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(pasta)
}
