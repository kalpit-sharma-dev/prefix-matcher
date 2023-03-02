package server

import (
	"log"
	"net/http"
	"prefix-matcher/src/handlers"
	"prefix-matcher/src/repository"
	"prefix-matcher/src/service"

	"github.com/gorilla/mux"
)

func LoadRoute() {
	log.Println("INFO : Loading Router")
	router := mux.NewRouter().PathPrefix("/prefix-matcher-service/api").Subrouter()
	router.Use(headerMiddleware)
	registerAppRoutes(router)
	log.Println("INFO : Router Loaded Successfully")
	log.Println("INFO : Application is started Successfully")
	http.ListenAndServe(":9999", router)
}

func registerAppRoutes(r *mux.Router) {
	log.Println("INFO : Registering Router ")

	trieRepo := repository.NewDatabaseRepository(repository.Data)

	eventService := service.NewService(trieRepo)

	prefixHandlers := handlers.NewHandler(eventService)

	r.HandleFunc("/search/prefix/{word}", prefixHandlers.HandleSearchWord).Methods(http.MethodGet)

	log.Println("INFO : Router Registered Successfully")
}

func headerMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Content-Type", "application/json; charset=utf-8")
		next.ServeHTTP(w, r)
	})
}
