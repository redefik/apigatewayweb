package main

import (
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/redefik/apigatewayweb/config"
	"github.com/redefik/apigatewayweb/microservice"
	"log"
	"net/http"
)

// healthCheck handles the requests coming from an external component responsible for verifying the status of the api
// gateway
func healthCheck(w http.ResponseWriter, _ *http.Request) {
	w.WriteHeader(http.StatusOK)
}

func main() {

	// Read the listening address of the gateway and the address of the other microservices
	err := config.SetConfigurationFromEnvironment()
	if err != nil {
		log.Panicln(err)
	}
	r := mux.NewRouter()
	// Register the handlers for the various HTTP requests
	r.HandleFunc("/didattica-mobile/api/v1.0/token", microservice.LoginUser).Methods(http.MethodPost)
	r.HandleFunc("/didattica-mobile/api/v1.0/courses/{by}/{string}", microservice.FindCourse).Methods(http.MethodGet)
	r.HandleFunc("/didattica-mobile/api/v1.0/teachingMaterials/{courseId}", microservice.FindCourseMaterials).Methods(http.MethodGet)
	r.HandleFunc("/didattica-mobile/api/v1.0/teachingMaterials/uploadLink/{filename}/course/{courseId}", microservice.GenerateTemporaryUploadLink).Methods(http.MethodGet)
	r.HandleFunc("/didattica-mobile/api/v1.0/teachingMaterials/downloadLink/{filename}/course/{courseId}", microservice.GenerateTemporaryDownloadLink).Methods(http.MethodGet)
	r.HandleFunc("/", healthCheck).Methods(http.MethodGet)
	// Wait for incoming requests. A new goroutine is created to serve each request

	// To make the API accessible from JavaScript
	originsOk := handlers.AllowedOrigins([]string{"*"})
	methodsOk := handlers.AllowedMethods([]string{"GET", "HEAD", "POST", "PUT", "OPTIONS"})
	headersOk := handlers.AllowedHeaders([]string{"Accept", "Accept-Language", "Content-Type", "Set-Cookie", "Authorization"})
	log.Fatal(http.ListenAndServe(config.Configuration.ApiGatewayAddress, handlers.CORS(originsOk, methodsOk, headersOk)(r)))
}
