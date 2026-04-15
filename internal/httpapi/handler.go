package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"

	"top-backend-test/go-postgres/internal/inventory"
)

type Handler struct {
	service *inventory.Service
}

func NewHandler(service *inventory.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) Routes() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", h.health)
	mux.HandleFunc("/api/inventory-items", h.inventoryCollection)
	mux.HandleFunc("/api/inventory-items/", h.inventoryMember)

	return mux
}

func (h *Handler) health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		methodNotAllowed(w)
		return
	}

	writeJSON(w, http.StatusOK, map[string]any{
		"success": true,
		"data": map[string]any{
			"status":  "ok",
			"service": "go-postgres",
		},
	})
}

func (h *Handler) inventoryCollection(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		items, err := h.service.List()
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]any{"success": false, "message": err.Error()})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{"success": true, "data": items})
	case http.MethodPost:
		var payload inventory.Payload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "message": "invalid json payload"})
			return
		}

		item, validationErrors, err := h.service.Create(payload)
		if len(validationErrors) > 0 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
				"success": false,
				"message": "validation failed",
				"errors":  validationErrors,
			})
			return
		}
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]any{"success": false, "message": err.Error()})
			return
		}

		writeJSON(w, http.StatusCreated, map[string]any{"success": true, "data": item})
	default:
		methodNotAllowed(w)
	}
}

func (h *Handler) inventoryMember(w http.ResponseWriter, r *http.Request) {
	id, err := parseID(r.URL.Path, "/api/inventory-items/")
	if err != nil {
		writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "message": "invalid inventory item id"})
		return
	}

	switch r.Method {
	case http.MethodGet:
		item, err := h.service.Find(id)
		if err != nil {
			writeJSON(w, http.StatusInternalServerError, map[string]any{"success": false, "message": err.Error()})
			return
		}
		if item == nil {
			writeJSON(w, http.StatusNotFound, map[string]any{"success": false, "message": "inventory item not found"})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{"success": true, "data": item})
	case http.MethodPut:
		var payload inventory.Payload
		if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "message": "invalid json payload"})
			return
		}

		item, validationErrors, err := h.service.Update(id, payload)
		if len(validationErrors) > 0 {
			writeJSON(w, http.StatusUnprocessableEntity, map[string]any{
				"success": false,
				"message": "validation failed",
				"errors":  validationErrors,
			})
			return
		}
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "message": err.Error()})
			return
		}
		if item == nil {
			writeJSON(w, http.StatusNotFound, map[string]any{"success": false, "message": "inventory item not found"})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{"success": true, "data": item})
	case http.MethodDelete:
		deleted, err := h.service.Delete(id)
		if err != nil {
			writeJSON(w, http.StatusBadRequest, map[string]any{"success": false, "message": err.Error()})
			return
		}
		if !deleted {
			writeJSON(w, http.StatusNotFound, map[string]any{"success": false, "message": "inventory item not found"})
			return
		}

		writeJSON(w, http.StatusOK, map[string]any{
			"success": true,
			"data":    map[string]any{"message": "inventory item deleted successfully"},
		})
	default:
		methodNotAllowed(w)
	}
}

func parseID(path, prefix string) (int64, error) {
	rawID := strings.TrimPrefix(path, prefix)
	return strconv.ParseInt(rawID, 10, 64)
}

func methodNotAllowed(w http.ResponseWriter) {
	writeJSON(w, http.StatusMethodNotAllowed, map[string]any{
		"success": false,
		"message": "method not allowed",
	})
}

func writeJSON(w http.ResponseWriter, status int, payload map[string]any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(payload)
}
