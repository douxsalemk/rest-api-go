package routes

import (
	"log"
	"net/http"
	"rest-api-go/controllers"
)



func HandleResquest() {
	http.HandleFunc("/", controllers.Home)
	http.HandleFunc("/api/personalidades", controllers.TodasPersonalidades)
	log.Fatal(http.ListenAndServe(":8000", nil))
}