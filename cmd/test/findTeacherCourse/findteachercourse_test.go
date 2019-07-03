package findTeacherCourse

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

// createTestGatewayFindTeacherCourse creates an http handler that handles the test requests
func createTestGatewayFindTeacherCourse() http.Handler {
	r := mux.NewRouter()
	r.HandleFunc("/didattica-mobile/api/v1.0/courses/{by}/{string}", microservice.FindCourse).Methods(http.MethodGet)
	return r
}

// TestFindCourseSuccess tests the following scenario: the client sends a course research request to the api gateway,
// passing the type of research (name or teacher) and the string representing the sequence to find.
// The gateway makes an http GET request with the given information and sends it to the user management microservice.
// It is assumed that exist a course with teacher matching with sequence "seq", so in this case the research has success.
func TestFindCourseSuccess(t *testing.T) {

	_ = config.SetConfigurationFromFile("../../../config/config-test.json")

	// generate a token to be appended to the request
	user := microservice.User{Name: "nome", Surname: "cognome", Username: "username", Password: "password", Type: "teacher"}
	token, _ := microservice.GenerateAccessToken(user, []byte(config.Configuration.TokenPrivateKey))

	// Make the get request for course searching
	request, _ := http.NewRequest(http.MethodGet, "/didattica-mobile/api/v1.0/courses/teacher/seq", nil)
	bearer := "Bearer " + token
	request.Header.Add("Authorization", bearer)

	response := httptest.NewRecorder()
	handler := createTestGatewayFindTeacherCourse()
	// a goroutine representing the microservice listens to the requests coming from the api gateway
	go mock.LaunchCourseManagementMock()
	// simulates a request-response interaction between client and api gateway
	handler.ServeHTTP(response, request)

	if response.Code != http.StatusOK {
		t.Error("Expected 200 OK but got " + strconv.Itoa(response.Code) + " " + http.StatusText(response.Code))
	}

}
