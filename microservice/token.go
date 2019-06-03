package microservice

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/redefik/apigatewayweb/config"
	"strings"

	"log"
	"net/http"
	"time"
)


// Claims encapsulates the payload that will be encoded to build the access token
// according to JWT standard
type Claims struct {
	Name string
	Surname string
	Type string
	jwt.StandardClaims
}

var ExpiredToken = errors.New("expired token")

// makeClaims builds the payload of the JWT token. The payload will contain:
// - user Name
// - user Surname
// - user Type (student or teacher)
// - token expiration time
func makeClaims(user User) (Claims) {
	expirationTime := time.Now().Add(time.Duration(10 * time.Minute)).Unix()
	claims := Claims{
		Name: user.Name,
		Surname: user.Surname,
		Type: user.Type,
		StandardClaims: jwt.StandardClaims{ExpiresAt: expirationTime},
	}
	return claims
}

// GenerateAccessToken builds the JWT token to be returned to the client.
func GenerateAccessToken(user User, signingKey []byte) (string, error) {
	claims := makeClaims(user)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(signingKey)
}

// ValidateToken parses the given token verifying that it has been signed with the proper key. It returns
// a Claims struct containing token payload and an error, not nil if the validation goes wrong
func ValidateToken(tokenString string, w http.ResponseWriter) (Claims, error) {
	claims := Claims{}
	jwtKey := []byte(config.Configuration.TokenPrivateKey)

	// The token string is parsed, decoded and stored into the given Claims struct
	token, err := jwt.ParseWithClaims(tokenString, &claims,
		func(token *jwt.Token) (interface{}, error) {
			return jwtKey, nil
		})

	// Check if the token has expired according to the expiry time fixed during the sign in
	if !token.Valid {
		err = ExpiredToken
		MakeErrorResponse(w, http.StatusUnauthorized, err.Error())
		log.Println(err.Error())
		return claims, err
	}

	// Check if the token has been signed with the private key of the api gateway
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			// If the token is expired or has not been signed according to the api gateway key, an Unauthorization code
			// is returned in both cases, but a different message is provided to the client.
			MakeErrorResponse(w, http.StatusUnauthorized, "Wrong credentials")
			log.Println("Wrong credentials")
			return claims, err
		}

		MakeErrorResponse(w, http.StatusBadRequest, "Malformed token")
		log.Println("Malformed token")
		return claims, err
	}

	return claims, nil

}


/*GetToken return the string representing the token inserted in the field Authorization Bearer of the header of http request.
  If token is not present or there are other problem in its retrieval an error response is sent to client*/
func GetToken(w http.ResponseWriter, r *http.Request) (string, error){

	authHeader := r.Header.Get("Authorization")
	// The token is provided in the form: Authorization: Bearer <Token>
	// Therefore a split is necessary
	authHeaderFields := strings.Split(authHeader, "Bearer")
	if len(authHeaderFields) != 2 {
		MakeErrorResponse(w, http.StatusBadRequest, "Bad request")
		log.Println("Couldn't parse the JWT token")
		return "", errors.New("error in parsing JWT token")
	}

	tokenString := strings.TrimSpace(authHeaderFields[1])

	return tokenString, nil
}
