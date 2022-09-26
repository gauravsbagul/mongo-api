package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gauravsbagul/mongo-api/router"
)

//*INFO: POSTMAN API LINK :https://www.getpostman.com/collections/ab8c9976f673a04c4612

func main() {

	fmt.Println("MongoDB API")

	r := router.Router()

	fmt.Println("Server is getting started...")

	log.Fatal(http.ListenAndServe(":4000", r))

	fmt.Println("Listening at post 4000...")

}
