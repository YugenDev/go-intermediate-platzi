package main

import "fmt"

// Las funciones variadicas en go pueden recibir como parametro una variable
// que represente un slice, lo que signififca que podemos pasar
// cualquier cantidad de argumentos a la hora de llamar a la función
// y eso que le pasemos se va a introducir a ese slice.
func sum(variables ...int) int {

	total := 0
	for _, num := range variables {
		total += num
	}

	return total
}

func printNames(names ...string){
	for _, name := range names {
		fmt.Println(name)
	}
}

func main() {
	fmt.Println(sum(1))
	fmt.Println(sum(2, 2))
	fmt.Println(sum(1, 2, 3, 4, 5, 9000))
	fmt.Println(sum(1, 1000, 9000))
	fmt.Println("")
	printNames("urato", "pirlin", "salo", "vic")
}
