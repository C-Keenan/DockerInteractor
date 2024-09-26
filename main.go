package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/C-Keenan/DockerInteractor/controllers"
	"github.com/joho/godotenv"
)

func main() {
	http.HandleFunc("/", controllers.ListAllContainers)
	err := godotenv.Load()
	if err != nil {
		fmt.Println("Error loading .env file")
		fmt.Println(err)
	}
	port := os.Getenv("PORT")
	fmt.Printf("Server is running on localhost%s\n", port)
	log.Fatal(http.ListenAndServe(port, nil))
}
