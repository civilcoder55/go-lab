package routers

import (
	"github.com/civilcoder55/go-lab/controllers"
	"github.com/gorilla/mux"
)

func GameRouter() *mux.Router {
	r := mux.NewRouter()
	r.HandleFunc("/", controllers.HomeHandler).Methods("GET")
	r.HandleFunc("/games", controllers.GetAllGamesHandler).Methods("GET")
	r.HandleFunc("/games/{id}", controllers.GetOneGameHandler).Methods("GET")
	r.HandleFunc("/games", controllers.CreateGameHandler).Methods("POST")
	r.HandleFunc("/games/{id}", controllers.UpdateGameHandler).Methods("PUT")
	r.HandleFunc("/games/{id}", controllers.DeleteGameHandler).Methods("DELETE")

	return r
}
