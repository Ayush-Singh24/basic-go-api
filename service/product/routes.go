package product

import (
	"fmt"
	"net/http"

	"github.com/Ayush-Singh24/basic-go-api/types"
	"github.com/Ayush-Singh24/basic-go-api/utils"
	"github.com/go-playground/validator"
	"github.com/gorilla/mux"
)

type Handler struct {
	store types.ProductStore
}

func NewHandler(store types.ProductStore) *Handler {
	return &Handler{store: store}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/products", h.handleGetProducts).Methods(http.MethodGet)
	router.HandleFunc("/products", h.handleCreateProduct).Methods(http.MethodPost)
}

func (h *Handler) handleGetProducts(w http.ResponseWriter, r *http.Request) {
	ps, err := h.store.GetProducts()
	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
	}

	utils.WriteJSON(w, http.StatusOK, ps)
}

func (h *Handler) handleCreateProduct(w http.ResponseWriter, r *http.Request) {
	//parse payload
	var payload types.ProductPayload
	if err := utils.ParseJSON(r, &payload); err != nil {
		utils.WriteError(w, http.StatusBadRequest, err)
	}

	// validate the payload
	if err := utils.Validate.Struct(payload); err != nil {
		errors := err.(validator.ValidationErrors)
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", errors))
		return
	}

	floatPrice, err := payload.Price.Float64()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}
	intQuantity, err := payload.Quantity.Int64()

	if err != nil {
		utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid payload %v", err))
		return
	}

	err = h.store.CreateProduct(types.Product{
		Name:        payload.Name,
		Description: payload.Description,
		Image:       payload.Image,
		Quantity:    int(intQuantity),
		Price:       floatPrice,
	})

	if err != nil {
		utils.WriteError(w, http.StatusInternalServerError, err)
		return
	}

	utils.WriteJSON(w, http.StatusCreated, nil)
}
