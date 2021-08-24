package httpHelper

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type HttpHandlerRepo struct {
}

func (h *HttpHandlerRepo) OpenServer(router *mux.Router, port string) {
	fmt.Println("server runs in port : ", port)
	log.Fatal(http.ListenAndServe(":"+port, router))
}

func (h *HttpHandlerRepo) HandleHttpRequest(req http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		req(w, r)
	}
}
func (h *HttpHandlerRepo) SendResponseMessage(w http.ResponseWriter, JSONObj interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "*")
	sendJSON, _ := json.Marshal(JSONObj)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write(sendJSON)
	if err != nil {
		return
	}
}

