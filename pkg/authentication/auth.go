package authentication

import (
	"encoding/json"
	"fmt"
	"github.com/golang-jwt/jwt"
	"net/http"
)
type ApiMessage struct {
	MessageAPI string
}
var authKey = []byte("User key")

func IsAuthorized(endpoint func(http.ResponseWriter, *http.Request)) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		tokenStr := r.Header["Jwt-Token"]
		if tokenStr != nil {
			token, err := jwt.Parse(tokenStr[0], func(token *jwt.Token) (interface{}, error) {

				if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
					return nil, fmt.Errorf("There  is an error")
				}
				return authKey, nil
			})
			fmt.Println("token err -> ", err)
			if err != nil {
				SendError(w, "Not Authorized")
			}

			if token.Valid {
				endpoint(w, r)
			}
		} else {
			SendError(w, "Not Authorized")
		}
	})
}

//SendError status for 401 StatusUnauthorized
func SendError(w http.ResponseWriter, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	apiMessage := ApiMessage{}
	apiMessage.MessageAPI = message
	w.WriteHeader(http.StatusUnauthorized)
	sendJSON, _ := json.Marshal(apiMessage)
	_, err := w.Write(sendJSON)
	if err != nil {
		fmt.Println(err)
		return
	}
}
