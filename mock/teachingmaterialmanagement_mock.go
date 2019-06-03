package mock

import (
	"github.com/gorilla/mux"
	"github.com/redefik/apigatewayweb/config"
	"net/http"
)

// simulates the behaviour of the teaching materials management microservice when it is asked for listing the materials
// available for a course
func TeachingMaterialsManagementMockFindCourseMaterials(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

// simulates the behaviour of the teaching materials management microservice when it is asked for generating a link
// for the upload of a new course material
func TeachingMaterialsManagementMockGenerateUploadLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	if filename == "courseId_file.txt" {
		w.WriteHeader(http.StatusOK)
	} else if filename == "courseId_existentFile.txt" {
		w.WriteHeader(http.StatusConflict)
	}
}

// simulates the behaviour of the teaching materials management microservice when it is asked for generating a link
// for the download of a course teaching material
func TeachingMaterialsManagementMockGenerateDownloadLink(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	filename := vars["filename"]
	if filename == "courseId_file.txt" {
		w.WriteHeader(http.StatusOK)
	} else if filename == "courseId_notExistentFile.txt" {
		w.WriteHeader(http.StatusNotFound)
	}
}

// starts a teaching materials microservice mock
func LaunchTeachingMaterialManagementMock() {
	r := mux.NewRouter()
	r.HandleFunc("/teaching_material_management/api/v1.0/list/{prefix}", TeachingMaterialsManagementMockFindCourseMaterials).Methods(http.MethodGet)
	r.HandleFunc("/teaching_material_management/api/v1.0/upload/{filename}", TeachingMaterialsManagementMockGenerateUploadLink).Methods(http.MethodGet)
	r.HandleFunc("/teaching_material_management/api/v1.0/download/{filename}", TeachingMaterialsManagementMockGenerateDownloadLink).Methods(http.MethodGet)
	http.ListenAndServe(config.Configuration.ApiGatewayAddress, r)
}
