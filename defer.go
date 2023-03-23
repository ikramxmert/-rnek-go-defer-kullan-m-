package main

import "fmt"

//Fonksiyonlar defer
var isConnected bool = false

func main() {
	fmt.Printf("Connection open : %v\n", isConnected)
	databaseProsseing()
	fmt.Printf("Connection open : %v\n", isConnected)
}

func databaseProsseing() {
	connect()
	fmt.Println("Defering disconnect!")
	defer disconnet() //buradaki işlemlerin en sonunda çağırıyor
	fmt.Printf("Connection open : %v\n", isConnected)
	fmt.Println("Doing Something")
}

func connect() {
	isConnected = true
	fmt.Println("Cennected to database!")
}

func disconnet() {
	isConnected = false
	fmt.Println("Disconnected")
}
