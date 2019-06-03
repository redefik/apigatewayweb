package microservice

import (
	"encoding/json"
	"errors"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"github.com/redefik/apigatewayweb/config"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
)

// hasPermissionOnCourse returns true if the given teacher holds the course with the provided identifier, false otherwise.
func hasPermissionOnCourse(teacherName string, teacherSurname, courseId string) (bool, error) {
	getResponse, err := http.Get(config.Configuration.CourseManagementAddress + "courses/teacher/" + teacherName + "-" + teacherSurname)
	if err != nil {
		return false, err
	}
	log.Println("Response status Code from Microservice: " + strconv.Itoa(getResponse.StatusCode))
	defer getResponse.Body.Close()
	if getResponse.StatusCode != http.StatusOK {
		return false, errors.New("internal error")
	}
	var getResponseBody []CourseMinified

	// Decode the course management microservice response
	jsonDecoder := json.NewDecoder(getResponse.Body)
	err = jsonDecoder.Decode(&getResponseBody)
	if err != nil {
		return false, err
	}
	courseHeldByTeacher := false
	for i := 0; i < len(getResponseBody); i++ {
		if courseId == getResponseBody[i].Id {
			courseHeldByTeacher = true
			break
		}
	}
	return courseHeldByTeacher, nil
}

// FindCourseMaterials process the request of teaching materials provided by a teacher for a given course. Two conditions are checked:
// 1. The request is properly authenticated
// 2. The teacher of the course is the user that made the request
func FindCourseMaterials(w http.ResponseWriter, r *http.Request) {
	/* For authentication purpose the access token is read from the Authorization Bearer header */
	tokenString, err := GetToken(w, r)
	if err != nil {
		return
	}
	/* The token is decoded and the claims are obtained for further checks */
	decodedToken , err := ValidateToken(tokenString, w)
	if err != nil {
		return
	}
	vars := mux.Vars(r) // url-encoded parameters
	courseId := vars["courseId"]
	teacherName := decodedToken.Name
	teacherSurname := decodedToken.Surname
	heldCourse, err := hasPermissionOnCourse(teacherName, teacherSurname, courseId)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}
	if !heldCourse {
		MakeErrorResponse(w, http.StatusUnauthorized, "Not Held Course")
		log.Println("Not Held Course")
		return
	}
	/* Upon successful validation, the client request is forwarded to the teaching materials management microservice and the response is
	returned to the client*/
	err = ForwardAndReturnGet(config.Configuration.TeachingMaterialManagementAddress + "list/" + courseId, w)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}

}

// GenerateTemporaryUploadLink process the request that the client makes to obtain a temporary link to upload the file
// with the provided name. The Api Gateway forwards the request to the teacher management microservice
// only if the teacher holds the course with the given identifier.
// The response of the microservice is converted in JSON to make it readable for the AngularJS client.
func GenerateTemporaryUploadLink(w http.ResponseWriter, r *http.Request) {
	/* For authentication purpose the access token is read from the Authorization Bearer header */
	tokenString, err := GetToken(w, r)
	if err != nil {
		return
	}
	/* The token is decoded and the claims are obtained for further checks */
	decodedToken, err := ValidateToken(tokenString, w)
	if err != nil {
		return
	}
	vars := mux.Vars(r) // url-encoded parameters
	courseId := vars["courseId"]
	filename := vars["filename"]
	teacherName := decodedToken.Name
	teacherSurname := decodedToken.Surname
	heldCourse, err := hasPermissionOnCourse(teacherName, teacherSurname, courseId)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}
	if !heldCourse {
		MakeErrorResponse(w, http.StatusUnauthorized, "Not Held Course")
		log.Println("Not Held Course")
		return
	}

	/* Upon successful validation, the client request is forwarded to the teaching materials management microservice and the response is
	returned to the client in JSON format*/
	getResponse, err := http.Get(config.Configuration.TeachingMaterialManagementAddress + "upload/" + courseId + "_" + filename)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}
	log.Println("Response status Code from Microservice: " + strconv.Itoa(getResponse.StatusCode))
	defer getResponse.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	if getResponse.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		getResponseBody, err := ioutil.ReadAll(getResponse.Body)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
		response := simplejson.New()
		response.Set("link", string(getResponseBody))
		responsePayload, err := response.MarshalJSON()
		if err != nil {
			log.Panicln(err)
		}
		w.Write(responsePayload)

	} else {
		w.WriteHeader(getResponse.StatusCode)
		responseBody, err := ioutil.ReadAll(getResponse.Body)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
		_, err = w.Write(responseBody)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
	}
}


// GenerateTemporaryDownloadLink process the request that the client makes to obtain a temporary link to download the file
// with the provided name. The Api Gateway forwards the request to the teacher management microservice
// only if the teacher holds the course with the given identifier.
// The response of the microservice is converted in JSON to make it readable from the AngularJS client.
func GenerateTemporaryDownloadLink(w http.ResponseWriter, r *http.Request) {
	/* For authentication purpose the access token is read from the Authorization Bearer header */
	tokenString, err := GetToken(w, r)
	if err != nil {
		return
	}
	/* The token is decoded and the claims are obtained for further checks */
	decodedToken , err := ValidateToken(tokenString, w)
	if err != nil {
		return
	}
	vars := mux.Vars(r) // url-encoded parameters
	courseId := vars["courseId"]
	filename := vars["filename"]
	teacherName := decodedToken.Name
	teacherSurname := decodedToken.Surname
	heldCourse, err := hasPermissionOnCourse(teacherName, teacherSurname, courseId)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}
	if !heldCourse {
		MakeErrorResponse(w, http.StatusUnauthorized, "Not Held Course")
		log.Println("Not Held Course")
		return
	}
	/* Upon successful validation, the client request is forwarded to the teaching materials management microservice and the response is
	returned to the client in JSON format*/
	getResponse, err := http.Get(config.Configuration.TeachingMaterialManagementAddress + "download/" + courseId + "_" + filename)
	if err != nil {
		MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
		log.Println("Api Gateway - Internal Server Error")
		return
	}
	log.Println("Response status Code from Microservice: " + strconv.Itoa(getResponse.StatusCode))
	defer getResponse.Body.Close()
	w.Header().Set("Content-Type", "application/json")
	if getResponse.StatusCode == http.StatusOK {
		w.WriteHeader(http.StatusOK)
		getResponseBody, err := ioutil.ReadAll(getResponse.Body)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
		response := simplejson.New()
		response.Set("link", string(getResponseBody))
		responsePayload, err := response.MarshalJSON()
		if err != nil {
			log.Panicln(err)
		}
		w.Write(responsePayload)

	} else {
		w.WriteHeader(getResponse.StatusCode)
		responseBody, err := ioutil.ReadAll(getResponse.Body)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
		_, err = w.Write(responseBody)
		if err != nil {
			MakeErrorResponse(w, http.StatusInternalServerError, "Api Gateway - Internal Server Error")
			log.Println("Api Gateway - Internal Server Error")
			return
		}
	}
}
