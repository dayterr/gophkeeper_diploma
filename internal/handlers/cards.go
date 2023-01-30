package handlers

import "net/http"

func (ah *AsyncHandler) PostCard(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (ah *AsyncHandler) ListCards(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}

func (ah *AsyncHandler) DeleteCard(w http.ResponseWriter, r *http.Request) {

	w.WriteHeader(http.StatusOK)
}