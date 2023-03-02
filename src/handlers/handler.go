package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"prefix-matcher/src/models"
	"prefix-matcher/src/service"

	"github.com/gorilla/mux"
)

type Handler struct {
	service service.IService
}

func NewHandler(service service.IService) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) HandleSearchWord(w http.ResponseWriter, r *http.Request) {
	log.Println("INFO : SearchWord")
	searchWord := mux.Vars(r)["word"]

	output, err := h.service.SearchWord(searchWord)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		failResponse := models.Error{
			Id:      http.StatusNotFound,
			Message: fmt.Sprintf("%v", err),
		}
		json.NewEncoder(w).Encode(failResponse)
		return
	}
	resp := models.Response{}
	w.WriteHeader(http.StatusOK)
	resp.LCP = output
	json.NewEncoder(w).Encode(&resp)
	return

}
