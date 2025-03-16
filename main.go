package main

import (
	"fmt"
	"golang-restaurant-order/app"
	"net/http"
)

func main() {
	db, _ := app.NewDB()

	server := http.Server{
		Addr: "localhost:8080",
	}

	fmt.Println("DB >>>", db)
	fmt.Printf("Server starting on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		fmt.Println(err.Error())
	}
}
