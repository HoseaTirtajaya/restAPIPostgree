package router

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/arganaphangquestian/gobasic/model"
	"github.com/gorilla/mux"
)

//Middleware
func apiMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		next.ServeHTTP(w, r)
	})
}

//Routing Handler
func handleBase(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(model.ResponseAPI{
		Status:  "Success",
		Message: "Hehehehehe",
	})
}

//Init router
func Init() {
	//PORT
	const PORT string = ":8080"

	myRouter := mux.NewRouter()
	myRouter.Use(apiMiddleware)
	myRouter.HandleFunc("/", handleBase).Methods("GET")
	myRouter.HandleFunc("/users", getAllUsers).Methods("GET")
	myRouter.HandleFunc("/register", registerHandler).Methods("POST")
	myRouter.HandleFunc("/login", loginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(PORT, myRouter))
}
