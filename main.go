package brilock

import (
	//"FirstService/handlers"
	"fmt"
	"log"
	"net/http"
)

func main() {
	log.Print("Starting the service...")
	router := handlers.Router()
	http.HandleFunc("/home", httpHandler)

	log.Print("The service is ready to listen and serve.")
	log.Fatal(http.ListenAndServe(":8000", router))
}

func httpHandler(w http.ResponseWriter, _ *http.Request) {
	fmt.Fprint(w, "Hello! Your request was processed.")
}
