package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	//http.HandleFunc("/ctrlist", controllers.ListAllContainers)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	fmt.Printf("Server is running on localhost:%s\n", port)
	http.ListenAndServe(port, nil)
}
