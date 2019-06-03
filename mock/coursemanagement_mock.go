package mock

import (
	"encoding/json"
	"github.com/bitly/go-simplejson"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

// The following type is used to simulate the response that the microservice
// gives when it is asked for the courses held by a teacher by the Api Gateway
// in the case of teaching material listing
type course struct {
	Id   string `json:"id"`
	Name string `json:"name"`
}

// CourseManagementMockSearchCourse simulates the behaviour of the course management microservice when receives a request of
// course research.
func CourseManagementMockSearchCourse(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	if mux.Vars(r)["by"] == "teacher" && mux.Vars(r)["string"] == "seq"{
		w.WriteHeader(http.StatusOK)
	// This is used when the material management use case is tested
	} else if mux.Vars(r)["by"] == "teacher" && mux.Vars(r)["string"] == "nome-cognome" {
		course1 := course{Id: "courseId", Name: "courseName1"}
		course2 := course{Id: "idCourse2", Name: "courseName2"}
		// The payload contains the courses held by the teaher "nome-cognome"
		responsePayload, err := json.Marshal([]course{course1, course2})
		if err != nil {
			log.Panicln(err)
		}
		_ , err = w.Write(responsePayload)
		if err != nil {
			log.Panic(err)
		}
		return
	} else if mux.Vars(r)["by"] == "notvalid" {
		w.WriteHeader(http.StatusBadRequest)
	} else {
		w.WriteHeader(http.StatusNotFound)
	}
	response := simplejson.New()
	response.Set("mock", "response")
	responsePayload, err := response.MarshalJSON()
	if err != nil {
		log.Panicln(err)
	}
	_ , err = w.Write(responsePayload)
	if err != nil {
		log.Panic(err)
	}
}

// starts a course management microservice mock
func LaunchCourseManagementMock() {
	r := mux.NewRouter()
	r.HandleFunc("/course_management/api/v1.0/courses/{by}/{string}", CourseManagementMockSearchCourse).Methods(http.MethodGet)
	http.ListenAndServe("0.0.0.0:81", r) // Listens on a different port to avoid conflict with teaching material mock
}
