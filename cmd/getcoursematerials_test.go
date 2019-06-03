package main

import (
"github.com/gorilla/mux"
"github.com/redefik/apigatewayweb/config"
"github.com/redefik/apigatewayweb/microservice"
"github.com/redefik/apigatewayweb/mock"
"net/http"
"net/http/httptest"
"strconv"
"testing"
)

// createTestGatewayGetCourseMaterials creates an http handler that handles the test requests
func createTestGatewayGetCourseMaterials() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/didattica-mobile/api/v1.0/teachingMaterials/{courseId}", microservice.FindCourseMaterials).Methods(http.MethodGet)
	return r
}

// TestFinCourseMaterialsSuccess tests the following scenario: the client requires the list of teaching material of a course
// and the microservice responds with 200 OK, so the client response shuold be 200 too.
func TestFindCourseMaterials(t *testing.T) {

	_ = config.SetConfiguration("../config/config-test.json")

	// generate a token to be appended to the request
	user := microservice.User{Name: "nome", Surname: "cognome", Username: "username", Password: "password", Type:"teacher"}
	token, _ := microservice.GenerateAccessToken(user, []byte(config.Configuration.TokenPrivateKey))

	// Make the get request for course searching
	request, _ := http.NewRequest(http.MethodGet, "/didattica-mobile/api/v1.0/teachingMaterials/courseId", nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	handler := createTestGatewayGetCourseMaterials()
	// a goroutine representing the course management microservice listens to the requests coming from the api gateway
	go mock.LaunchCourseManagementMock()
	// a goroutine representing the teachging material microservice listens to the requests coming from the api gateway
	go mock.LaunchTeachingMaterialManagementMock()
	// simulates a request-response interaction between client and api gateway
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Error("Expected 200 OK but got " + strconv.Itoa(response.Code) + " " + http.StatusText(response.Code))
	}
}
