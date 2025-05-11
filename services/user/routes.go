package user

import (
	"net/http"

	"github.com/HerbertSanches/ecommerce-go/types"
	"github.com/HerbertSanches/ecommerce-go/utils"
	"github.com/gorilla/mux"
)

type Handler struct {
	store *types.UserStore
}

func NewHandler() *Handler {
	return &Handler{}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("POST")
	router.HandleFunc("/register", h.handleRegister).Methods("POST")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) handleRegister(w http.ResponseWriter, r *http.Request) {
	//get json payload
	var payload types.RegisterUserPayload
	if err := utils.ParseJSON(r, payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	//check if the user exists


	//if it doesnt we create the new user
}
