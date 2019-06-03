package microservice

/* Here you can find the structures that encapsulates the fields of the JSON body belonging to the http requests and replies*/

// Encapsulates the fields of the JSON body of the http GET requests sent to the user management microservice from
// the api gateway
type LoginRequestBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	Name string 	 `json:"name"`
	Surname string   `json:"surname"`
	Username string  `json:"username"`
	Password string  `json:"password"`
	Type string      `json:"type"`
}

// Encapsulates the field of the JSON body of the http message sent by the user management microservice when it responds
// to a http get request.
type LoginResponseBody  struct {
	User User `json:"user"`
}

// Encapsulates the field of the JSON error response from a microservice
type ErrorResponse struct {
	Error string `json:"error"`
}

// Encapsulates the field of the JSON body of the http POST request that the Api Gateway makes in order to create a student
type StudentCreationRequest struct {
	Name string `json:"name"`
	Username string `json:"username"`
}

// Encapsulates the fields of the JSON body of the http message sent by the course management microservice when
// a reserach by teacher is performed. Notice that only the field id is parsed because it is the only used by the
// api gateway
type CourseMinified struct {
	Id string `json:"id"`
}
