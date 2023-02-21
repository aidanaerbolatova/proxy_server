package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"proxy/models"
)

func (h *Handler) ProxyServe(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
		return
	case http.MethodPost:
		var request models.Request
		if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
			http.Error(w, "error with decode response body", http.StatusInternalServerError)
			return
		}
		resp, err := h.proxy.ProxyRequest(request)
		if err != nil {
			fmt.Println(err)
			http.Error(w, "error with get response from proxy", http.StatusInternalServerError)
			return
		}
		w.WriteHeader(200)
		if err := json.NewEncoder(w).Encode(resp); err != nil {
			http.Error(w, "error convert response", http.StatusInternalServerError)
			return
		}
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

}
