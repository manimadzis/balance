package server

import (
	. "balance/internal/balance/entities"
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"strings"
)

func (s *Server) handleBalanceUpdate() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction SystemTransaction

		if err := parseSystemTransaction(r.Body, &transaction); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.SaveSystemTransaction(&transaction); err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		respondJSON(w, http.StatusAccepted, transaction)
	}
}

func (s *Server) handleTransaction() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var transaction Transaction

		if err := parseTransaction(r.Body, &transaction); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.SaveTransaction(&transaction); err != nil {
			respondError(w, http.StatusInternalServerError, ErrDbFault)
			return
		}

		respondJSON(w, http.StatusAccepted, transaction)
	}
}

func (s *Server) handleGetBalance() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var account Account

		values := r.URL.Query()

		if len(values["params"]) == 0 {
			respondError(w, http.StatusBadRequest, errors.New("Empty 'params'"))
			return
		}

		if err := parseAccount(strings.NewReader(values["params"][0]), &account); err != nil {
			respondError(w, http.StatusBadRequest, err)
			return
		}

		if err := s.store.GetBalance(&account); err != nil {
			respondError(w, http.StatusInternalServerError, err)
			return
		}

		respondJSON(w, http.StatusOK, account)
	}
}

func respondJSON(w http.ResponseWriter, status int, data interface{}) {

	jsonBytes, err := json.Marshal(data)
	if err != nil {
		log.Printf("Cannot marshal %v: %v", data, err)
	}

	respond(w, status, string(jsonBytes))
}

func respondError(w http.ResponseWriter, status int, err error) {

	type errorResponse struct {
		ErrMsg string `json:"error_msg"`
	}

	jsonBytes, err := json.Marshal(
		errorResponse{
			ErrMsg: err.Error(),
		})
	if err != nil {
		jsonBytes = []byte{}
	}

	respond(w, status, string(jsonBytes))
}

func respond(w http.ResponseWriter, status int, msg string) {
	w.WriteHeader(status)
	w.Write([]byte(msg))
}
