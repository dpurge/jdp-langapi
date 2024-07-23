package main

import (
	"fmt"
	"os"

	"github.com/dpurge/lang-api/pkg/server"
)

// func handleOther(w http.ResponseWriter, r *http.Request) {
// 	log.Println("Received a non domain request")
// 	w.Write([]byte("Hello, stranger..."))
// }

func main() {
	// router := http.NewServeMux()
	// router.HandleFunc("/", handleOther)

	// server := http.Server{
	// 	Addr:    ":8080",
	// 	Handler: router,
	// }

	// fmt.Println("Server listening on port :8080")
	// server.ListenAndServe()

	app := server.App{}
	app.Initialize(
		os.Getenv("APP_DB_USERNAME"),
		os.Getenv("APP_DB_PASSWORD"),
		os.Getenv("APP_DB_NAME"),
	)

	fmt.Println("Server listening on port :8080")
	app.Run(":8080")
}
