package handlers

import (
	"encoding/json"
	"github.com/farhodm/banner-service-clean-arch/internal/helpers"
	"github.com/farhodm/banner-service-clean-arch/internal/usecase"
	"github.com/go-playground/validator/v10"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

type BannerHandler struct {
	useCase   *usecase.BannerUseCase
	validator *validator.Validate
}

func NewBannerHandler(uc *usecase.BannerUseCase) *BannerHandler {
	return &BannerHandler{
		useCase:   uc,
		validator: validator.New(),
	}
}

type createBannerRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (h *BannerHandler) CreateBanner(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	var request createBannerRequest
	validationErrors := make(map[string]string)

	// read and decode request body to json
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&request); err != nil {
		helpers.ValidationsError(err, validationErrors)
		w.Header().Set("Content-Type", "application/json")
		_ = json.NewEncoder(w).Encode(validationErrors)
		return
	}

	banner, err := h.useCase.CreateBanner(r.Context(), request.Title, request.Content)
	if err != nil {
		http.Error(w, "Fail create banner", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(banner)
}

func (h *BannerHandler) GetAllBanners(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	banners := h.useCase.GetAllBanners(r.Context())

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	_ = json.NewEncoder(w).Encode(banners)
}

func (h *BannerHandler) GetBanner(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	bannerId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	banner, err := h.useCase.GetBanner(r.Context(), bannerId)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(banner)
}

type UpdateBannerRequest struct {
	Title   string `json:"title" validate:"required"`
	Content string `json:"content" validate:"required"`
}

func (h *BannerHandler) UpdateBanner(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	validationErrors := make(map[string]string)

	id := p.ByName("id")
	bannerId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	var request UpdateBannerRequest

	// read and decode request body to json
	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		http.Error(w, "Invalid JSON", http.StatusBadRequest)
		return
	}

	if err := h.validator.Struct(&request); err != nil {
		helpers.ValidationsError(err, validationErrors)
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(validationErrors)
		return
	}

	if err = h.useCase.UpdateBanner(r.Context(), bannerId, request.Title, request.Content); err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_ = json.NewEncoder(w).Encode(err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("message: Successfully updated"))
}

func (h *BannerHandler) DeleteBanner(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	id := p.ByName("id")

	bannerId, err := strconv.Atoi(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	if err = h.useCase.DeleteBanner(r.Context(), bannerId); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Write([]byte("Successfully deleted"))
}
