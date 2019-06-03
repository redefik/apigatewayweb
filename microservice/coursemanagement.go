package microservice

import (
	"github.com/gorilla/mux"
	"github.com/redefik/apigatewayweb/config"
	"log"
	"net/http"
)


// FindCourse process the course searching request coming from the client and validate the embedded access token. It verify
// if the token is properly signed and not expired. Upon successful validation, the request is forwarded to the microservice
// and the response is forwarded to the client.
func FindCourse(w http.ResponseWriter, r *http.Request) {

	/* For authentication purpose the access token is read from the Authorization Bearer header */
	tokenString, err := GetToken(w, r)
	if err != nil {
		return
	}
	/* The token is decoded and the claims are obtained for further checks */
	_ , err = ValidateToken(tokenString, w)
	if err != nil {
		return
	}
	/* Upon successful validation, the request is forwarded to the course management microservice and the response is
	returned to the client*/
	vars := mux.Vars(r) // url-encoded parameters
	by := vars["by"]
	searchString := vars["string"]
	err = ForwardAndReturnGet(config.Configuration.CourseManagementAddress + "courses" + "/" + by + "/" + searchString, w)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}

}
