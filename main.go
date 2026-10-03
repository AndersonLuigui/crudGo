package main

import (
	"fmt"
	"net/http"
)

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World!")

}

func main() {

	var num1 int
	var num2 int

	fmt.Println("Digite o primeiro número")
	fmt.Scanln(&num1)

	fmt.Println("Dígite o segundo número")
	fmt.Scanln(&num2)

	soma := num1 + num2

	fmt.Println("A soma dos números ", num1, " e ", num2, " é ", soma)

	http.HandleFunc("/hello", hello)
	fmt.Println("Server is running on port 8080")

	http.ListenAndServe(":8080", nil)

}
